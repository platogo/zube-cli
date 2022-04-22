package zube

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/platogo/zube-cli/zube/models"
)

const (
	ZubeHost  string = "zube.io"
	ApiUrl    string = "https://zube.io/api/"
	UserAgent string = "Zube-CLI"
)

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

// Encodes everything in `Query` into a flat Zube query string
func (query *Query) Encode() string {
	q := url.Values{}

	// Pagination
	q.Add("page", query.Pagination.Page)
	q.Add("per_page", query.Pagination.PerPage)
	// Order
	q.Add("order[by]", query.Order.By)
	q.Add("order[direction]", query.Order.Direction)

	// Filter
	for field, v := range query.Filter.Where {
		switch value := v.(type) {
		case []string:
			q.Add(fmt.Sprintf("where[%s][]", field), strings.Join(value, ","))
		default:
			q.Add(fmt.Sprintf("where[%s]", field), fmt.Sprint(v))
		}
	}

	for _, col := range query.Filter.Select {
		q.Add("select[]", col)
	}
	return q.Encode()
}

type Client struct {
	models.ZubeAccessToken // An encoded access JWT valid for 24h to the Zube API
	Host                   string
	ClientId               string // Your unique client ID
}

// Sets up a client with a profile, and caches it if needed
func NewClientWithProfile(profile *Profile) (*Client, error) {
	client := NewClient(profile.ClientId)

	if profile.IsTokenValid() {
		client.AccessToken = profile.AccessToken
	} else {
		// Refresh client token and dump it to profile
		privateKey, err := GetPrivateKey()
		if err != nil {
			log.Fatalln(err)
			return client, err
		}

		profile.AccessToken, err = client.RefreshAccessToken(privateKey)

		if err != nil {
			log.Fatalln(err)
			return client, err
		}

		ok := profile.SaveToConfig()

		if ok != nil {
			log.Fatal("Failed to save current configuration:", ok)
		}
	}
	return client, nil
}

// Constructs a new client with only host and Client ID configured, enough to make an access token request.
func NewClient(clientId string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId}
}

// Like `NewClient`, but requires and access token ready to be used for API requests.
func NewClientWithAccessToken(clientId, accessToken string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId, ZubeAccessToken: models.ZubeAccessToken{AccessToken: accessToken}}
}

// Fetch the access token JWT from Zube API and set it for the client. If it already exists, refresh it.
func (client *Client) RefreshAccessToken(key *rsa.PrivateKey) (string, error) {
	refreshJWT, err := GenerateRefreshJWT(client.ClientId, key)

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
	url := zubeURL("/api/current_person", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch current person info!")
	}

	json.Unmarshal(body, &currentPerson)
	return currentPerson
}

// Fetch and return an array of `Card`s
func (client *Client) FetchCards(query *Query) []models.Card {
	var response models.PaginatedResponse[models.Card]

	url := zubeURL("/api/cards", *query)

	// TODO: Support pagination
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch list of cards")
	}

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchCardComments(cardId int) []models.Comment {
	var response models.PaginatedResponse[models.Comment]

	url := zubeURL("/api/cards/"+fmt.Sprint(cardId)+"/comments", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatalf("Failed to fetch comments for card with Id: %d", cardId)
	}

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) CreateCard() {}

func (client *Client) FetchWorkspaces() []models.Workspace {
	var response models.PaginatedResponse[models.Workspace]

	url := zubeURL("/api/workspaces", Query{})
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch list of workspaces")
	}

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch and return an array of `Account`s
func (client *Client) FetchAccounts() []models.Account {
	var response models.PaginatedResponse[models.Account]

	url := zubeURL("/api/accounts", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch list of accounts")
	}

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchProjects() []models.Project {
	var response models.PaginatedResponse[models.Project]

	url := zubeURL("/api/projects", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatal("Failed to fetch list of projects")
	}

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch cards for a specific project. The `project_id` key in the `Where` part of the `Query`'s `Filter` will have no effect.
func (client *Client) FetchProjectCards(projectId int, query *Query) []models.Card {
	var response models.PaginatedResponse[models.Card]

	url := zubeURL(fmt.Sprintf("/api/projects/%d/cards", projectId), *query)

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	if err != nil {
		log.Fatalf("Failed to fetch cards for project with Id: %d", projectId)
	}

	json.Unmarshal(body, &response)
	return response.Data
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

func zubeURL(path string, query Query) url.URL {
	return url.URL{Scheme: "https", Host: ZubeHost, Path: path, RawQuery: query.Encode()}
}
