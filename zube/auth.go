package zube

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
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
