package zube

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/yaml.v3"
)

type Profile struct {
	ClientId    string `yaml:"client_id"`    // Populated when a profile is read from config file
	AccessToken string `yaml:"access_token"` // Populated when a profile is read from config file
}

func ConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "config", "zube")
}

func ConfigFile() string {
	return filepath.Join(ConfigDir(), "config.yml")
}

func (profile *Profile) SaveToConfig() error {
	data, err := yaml.Marshal(profile)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err2 := ioutil.WriteFile(ConfigFile(), data, 0)

	if err2 != nil {
		log.Fatal(err2)
		return err2
	}

	return nil
}

// Check if the locally saved Access Token JWT is still valid.
func (profile *Profile) IsAccessTokenExpired() (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(profile.AccessToken, jwt.MapClaims{})

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

// Checks if the profile's access token is present and not expired
func (profile *Profile) IsTokenValid() bool {
	if profile.AccessToken != "" {
		isExp, _ := profile.IsAccessTokenExpired()
		return !isExp
	}

	return false
}
