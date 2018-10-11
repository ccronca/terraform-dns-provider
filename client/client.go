package client

import (
	"net/http"
)

// Client is the object that handles talking to the DNS API. This maintains
// state information for a particular dns requests
type Client struct {
	apiUser, apiPassword, baseUrl string

	//The Http Client that is used to make requests
	HttpClient *http.Client
}

// NewClient returns a new client which can be used to access the API
// methods. The expected argument is the API user and password.
func NewClient(apiUser, apiPassword string, c *http.Client) *Client {

	return &Client{
		apiUser:     apiUser,
		apiPassword: apiPassword,
		baseUrl:     "",
		HttpClient:  c,
	}
}

// SetBaseUrl changes the value of baseUrl.
func (c *Client) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

// GetBaseUrl returns the baseUrl.
func (c *Client) GetBaseUrl() string {
	return c.baseUrl
}
