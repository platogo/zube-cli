package zube

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/platogo/zube-cli/zube/models"
)

const (
	ZubeHost string = "zube.io"
	ApiUrl   string = "https://zube.io/api/"
)

var UserAgent = "Zube-CLI"

// Request parameter struct definitions
type Pagination struct {
	Page    string // Results page to paginate to
	PerPage string // How many results to fetch per page, defaults to 30
}

type Order struct {
	By        string // Column / field to order by
	Direction string // Either of `asc` or `desc`
}

type Filter struct {
	Where  map[string]any // Map of keys corresponding to fields to filter by
	Select []string       // Array of attributes to select
}

// Represents possible Zube query parameters
type Query struct {
	Pagination
	Order
	Filter
}

type Client struct {
	models.ZubeAccessToken // An encoded access JWT valid for 24h to the Zube API
	Host                   string
	ClientId               string // Your unique client ID
}

func NewClient(clientId string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId}
}

func NewClientWithAccessToken(clientId, accessToken string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId, ZubeAccessToken: models.ZubeAccessToken{AccessToken: accessToken}}
}

// Fetch the access token JWT from Zube API and set it for the client. If it already exists, refresh it.
func (client *Client) RefreshAccessToken(key *rsa.PrivateKey) (string, error) {
	refreshJWT, err := generateRefreshJWT(client.ClientId, key)

	req, _ := zubeAccessTokenRequest(http.MethodPost, ApiUrl+"users/tokens", nil, client.ClientId, refreshJWT)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	data := models.ZubeAccessToken{}
	json.Unmarshal(body, &data)
	client.AccessToken = string(data.AccessToken)
	return client.AccessToken, err
}

func (client *Client) FetchCurrentPerson() models.CurrentPerson {
	currentPerson := models.CurrentPerson{}
	url := url.URL{Scheme: "https", Host: ZubeHost, Path: "/api/current_person"}

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch current person info!")
	}

	json.Unmarshal(body, &currentPerson)
	return currentPerson
}

func (client *Client) FetchCards() []models.Card {
	var cards models.Cards

	url := url.URL{Scheme: "https", Host: ZubeHost, Path: "/api/cards"}

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch list of cards")
	}

	json.Unmarshal(body, &cards)
	return cards.Data
}

// Wrapper around `performAPIRequestURL` for e.g. GET requests with no request body
func (client *Client) performAPIRequestURLNoBody(method string, url *url.URL) ([]byte, error) {
	return client.performAPIRequestURL(method, url, nil)
}

// Performs a generic request with URL and body
func (client *Client) performAPIRequestURL(method string, url *url.URL, body io.Reader) ([]byte, error) {
	req, _ := http.NewRequest(method, url.String(), body)

	if client.AccessToken == "" {
		return nil, errors.New("missing access token")
	}
	req.Header.Add("Authorization", "Bearer "+client.AccessToken)
	req.Header.Add("X-Client-ID", client.ClientId)
	req.Header.Add("User-Agent", UserAgent)
	if body != nil && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) {
		req.Header.Add("Accept", "application/json")
	}

	resp, _ := http.DefaultClient.Do(req)

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return respBody, err
}

// Only used to create a request to fetch an access token JWT using a refresh JWT
func zubeAccessTokenRequest(method, url string, body io.Reader, clientId, refreshJWT string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+refreshJWT)
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
