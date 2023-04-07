package zube

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/platogo/zube/models"
)

const PrivateKeyFileName = "zube_api_key.pem"

// Returns the correct path to the user's private key file
func PrivateKeyFilePath() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".ssh", PrivateKeyFileName)
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyFile, err := ioutil.ReadFile(PrivateKeyFilePath())
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// Create a refresh JWT valid for one minute, used to fetch an access token JWT
func GenerateRefreshJWT(clientId string, key *rsa.PrivateKey) (string, error) {
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

// Returns true if the token is present and not expired
func IsTokenValid(token models.ZubeAccessToken) bool {
	if token.AccessToken != "" {
		isExp, _ := isAccessTokenExpired(token.AccessToken)
		return !isExp
	}

	return false
}

// Returns true if the JWT is expired
func isAccessTokenExpired(accessToken string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})

	if err != nil {
		log.Fatal(err)
		return true, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Fatalf("Can't convert token's claims to standard claims")
	}

	var expTime time.Time
	now := time.Now()

	switch exp := claims["exp"].(type) {
	case float64:
		expTime = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		expTime = time.Unix(v, 0)
	}

	isExpired := expTime.Unix() < now.Unix()

	return isExpired, nil
}
