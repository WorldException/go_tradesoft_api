package tradesoft

import (
	"tradesoft/analog"
	"tradesoft/info"
	"tradesoft/messenger"
	"tradesoft/provider"

	"github.com/go-resty/resty/v2"
)

// Client represents the Tradesoft API client
type Client struct {
	baseURL    string
	httpClient *resty.Client
	provider   *provider.Service
	info       *info.Service
	analog     *analog.Service
	messenger  *messenger.Service
}

// NewClient creates a new instance of Tradesoft API client
func NewClient(baseURL string) *Client {
	httpClient := resty.New()

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		provider:   provider.NewService(httpClient, baseURL),
		info:       info.NewService(httpClient, baseURL),
		analog:     analog.NewService(httpClient, baseURL),
		messenger:  messenger.NewService(httpClient, baseURL),
	}
}

// SetAuth sets authentication credentials for the client
func (c *Client) SetAuth(user, password string) {
	c.httpClient.SetBasicAuth(user, password)
}

// Provider returns a provider service client
func (c *Client) Provider() *provider.Service {
	return c.provider
}

// Info returns an info service client
func (c *Client) Info() *info.Service {
	return c.info
}

// Analog returns an analog service client
func (c *Client) Analog() *analog.Service {
	return c.analog
}

// Messenger returns a messenger service client
func (c *Client) Messenger() *messenger.Service {
	return c.messenger
}
