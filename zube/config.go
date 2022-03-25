package zube

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func IsConfigFilePresent() bool {
	_, err := os.Stat(ConfigFile())
	return !errors.Is(err, os.ErrNotExist)
}

// Parses a locally saved profile and returns it
func ParseDefaultConfig() (Profile, error) {
	return parseConfigFile(ConfigFile())
}

// Check if the locally saved Access Token JWT is still valid.
func (c *Profile) IsAccessTokenExpired() (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(c.AccessToken, jwt.MapClaims{})

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

func parseConfigFile(filename string) (Profile, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalf(err.Error())
	}

	var profile Profile

	err = yaml.Unmarshal(data, &profile)

	if err != nil {
		log.Fatal(err)
		return profile, err
	}

	return profile, nil
}
