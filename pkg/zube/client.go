package zube

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/platogo/cache"
	"github.com/platogo/zube/models"
	"github.com/spf13/viper"
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
	Search string `json:"search"` // Undocumented search API query parameter, pass search query string here
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

	// Search
	if query.Search != "" {
		q.Add("search", query.Search)
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
	HTTPc                  http.Client
}

// Creates and returns a Zube Client with an access token
// If the current access token is invalid, it is refreshes and saved to config
func NewClient() (*Client, error) {
	client := &Client{
		ClientId:        viper.GetString("client_id"),
		ZubeAccessToken: models.ZubeAccessToken{AccessToken: viper.GetString("access_token")},
		HTTPc:           http.Client{Timeout: time.Duration(10) * time.Second},
	}

	if !IsTokenValid(client.ZubeAccessToken) {
		privateKey, err := GetPrivateKey()

		if err != nil {
			log.Fatalln(err)
			return client, err
		}
		access_token, err := client.RefreshAccessToken(privateKey)

		if err != nil {
			log.Fatalln(err)
			return client, err
		}

		viper.Set("access_token", access_token)
		viper.WriteConfig()
		client.ZubeAccessToken.AccessToken = access_token
	}

	return client, nil
}

// Constructs a new client with only host and Client ID configured, enough to make an access token request.
func NewClientWithId(clientId string) *Client {
	return &Client{Host: ZubeHost, ClientId: clientId}
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

	Check(err, "failed to fetch current person info!")

	json.Unmarshal(body, &currentPerson)
	return currentPerson
}

// Fetch and return an array of `Card`s
func (client *Client) FetchCards(query *Query) []models.Card {
	var response models.PaginatedResponse[models.Card]

	url := zubeURL("/api/cards", *query)

	// TODO: Support pagination
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, "failed to fetch cards!")

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchCardComments(cardId int) []models.Comment {
	var response models.PaginatedResponse[models.Comment]

	url := zubeURL("/api/cards/"+fmt.Sprint(cardId)+"/comments", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch comments for card with Id: %d", cardId))

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) CreateCard(card *models.Card) models.Card {
	var respCard models.Card
	url := zubeURL("/api/cards", Query{})
	data, _ := json.Marshal(card)
	resp, err := client.performAPIRequestURL(http.MethodPost, &url, bytes.NewBuffer(data))

	Check(err, "failed to create card!")

	json.Unmarshal(resp, &respCard)

	return respCard
}

// Search Zube cards using a simple Query struct with `search` field in it.
func (client *Client) SearchCards(query *Query) []models.Card {
	var response models.PaginatedResponse[models.Card]

	url := zubeURL("/api/cards", *query)
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to find card with text: %s", query.Search))

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchWorkspaces(query *Query) []models.Workspace {
	var response models.PaginatedResponse[models.Workspace]

	url := zubeURL("/api/workspaces", *query)
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, "failed to fetch workspaces!")

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch all epics for a given project
func (client *Client) FetchEpics(projectId int) []models.Epic {
	var response models.PaginatedResponse[models.Epic]

	url := zubeURL(fmt.Sprintf("/api/projects/%d/epics", projectId), Query{})
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch card for project with Id: %d", projectId))

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch and return an array of `Account`s
func (client *Client) FetchAccounts(query *Query) []models.Account {
	var response models.PaginatedResponse[models.Account]

	url := zubeURL("/api/accounts", *query)

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, "failed to fetch accounts")

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch and return an array of Github `Source`s
func (client *Client) FetchSources() []models.Source {
	var response models.PaginatedResponse[models.Source]

	url := zubeURL("/api/sources", Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, "failed to fetch sources")

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchProjects(query *Query) []models.Project {
	var response models.PaginatedResponse[models.Project]

	url := zubeURL("/api/projects", *query)

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, "failed to fetch projects")

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch cards for a specific project. The `project_id` key in the `Where` part of the `Query`'s `Filter` will have no effect.
func (client *Client) FetchProjectCards(projectId int, query *Query) []models.Card {
	var response models.PaginatedResponse[models.Card]

	url := zubeURL(fmt.Sprintf("/api/projects/%d/cards", projectId), *query)

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch cards for project with Id: %d", projectId))

	json.Unmarshal(body, &response)
	return response.Data
}

func (client *Client) FetchProjectMembers(projectId int) []models.Member {
	var response models.PaginatedResponse[models.Member]

	url := zubeURL(fmt.Sprintf("/api/projects/%d/members", projectId), Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch project members for project with Id: %d", projectId))

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch all labels for a given project
func (client *Client) FetchLabels(projectId int) []models.Label {
	var response models.PaginatedResponse[models.Label]

	url := zubeURL(fmt.Sprintf("/api/projects/%d/labels", projectId), Query{})

	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch labels for project with Id: %d", projectId))

	json.Unmarshal(body, &response)
	return response.Data
}

// Fetch all sprints for a given workspace
func (client *Client) FetchSprints(workspaceId int) []models.Sprint {
	var response models.PaginatedResponse[models.Sprint]

	url := zubeURL(fmt.Sprintf("/api/workspaces/%d/sprints", workspaceId), Query{})
	body, err := client.performAPIRequestURLNoBody(http.MethodGet, &url)

	Check(err, fmt.Sprintf("failed to fetch sprints for workspace with ID: %d", workspaceId))

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

	urlsha1 := sha1.Sum([]byte(url.String()))
	cacheKey := fmt.Sprintf("%x", urlsha1)
	cached, found := cache.Get(cacheKey)

	if found {
		req.Header.Add("If-None-Match", cached.Etag)
	}

	req.Header.Add("Authorization", "Bearer "+client.AccessToken)
	req.Header.Add("X-Client-ID", client.ClientId)
	req.Header.Add("User-Agent", UserAgent)
	if body != nil && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch) {
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
	}

	resp, _ := client.HTTPc.Do(req)

	// If cache exists and has not been changed on server, return cache data
	if found && resp.StatusCode == http.StatusNotModified {
		return json.Marshal(cached.Data)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if etag := resp.Header.Get("ETag"); etag != "" {
		cache.Save(cacheKey, etag, respBody)
	}

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
