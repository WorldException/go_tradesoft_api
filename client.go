package go_tradesoft_api

import (
	"net/http"

	"github.com/WorldException/go_tradesoft_api/analog"
	"github.com/WorldException/go_tradesoft_api/info"
	"github.com/WorldException/go_tradesoft_api/messenger"
	"github.com/WorldException/go_tradesoft_api/provider"

	"github.com/go-resty/resty/v2"
)

var (
	TradeSoftUrl = "https://service.tradesoft.ru/3"
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

func NewClientDefault() *Client {
	return NewClient(TradeSoftUrl)
}

func (c *Client) SetHttpClient(client *http.Client) {
	restClient := resty.NewWithClient(client)
	c.httpClient = restClient
	c.provider = provider.NewService(restClient, c.baseURL)
	c.info = info.NewService(restClient, c.baseURL)
	c.analog = analog.NewService(restClient, c.baseURL)
	c.messenger = messenger.NewService(restClient, c.baseURL)
}

// Авторизация для доступа к трейдсофт
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
