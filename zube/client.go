package zube

import (
	"crypto/rsa"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ZubeHost string = "zube.io"
	ApiUrl   string = "https://zube.io/api/"
)

var UserAgent = "Zube-CLI"

type ZubeAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Client struct {
	Host        string
	ClientId    string // Your unique client ID
	AccessToken string // An encoded access JWT valid for 24h to the Zube API
}

func NewClient(clientId string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId}
}

func NewClientWithAccessToken(host, clientId, accessToken string) *Client {
	return &Client{Host: host, ClientId: clientId, AccessToken: accessToken}
}

// Fetch the access token JWT from Zube API and set it for the client. If it already exists, refresh it.
func (client *Client) RefreshAccessToken(key *rsa.PrivateKey) (string, error) {
	refreshJWT, err := generateRefreshJWT(client.ClientId, key)

	req, _ := zubeRequest(http.MethodPost, ApiUrl+"users/tokens", nil, client.ClientId, refreshJWT)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	data := ZubeAccessTokenResponse{}
	json.Unmarshal(body, &data)
	client.AccessToken = string(data.AccessToken)
	return client.AccessToken, err
}

func zubeRequest(method, url string, body io.Reader, clientId, token string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-Client-ID", clientId)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Accept", "application/json")
	return req, nil
}

// Create a refresh JWT valid for one minute, used to fetch an access token JWT
func generateRefreshJWT(clientId string, key *rsa.PrivateKey) (string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: (now.Add(time.Minute)).Unix(),
		Issuer:    clientId,
	}
	if err := claims.Valid(); err != nil {
		log.Fatalf("invalid claims: %s", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}
