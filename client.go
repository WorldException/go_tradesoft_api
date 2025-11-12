package go_tradesoft_api

import (
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
	baseURL        string
	httpClient     *resty.Client
	provider       *provider.Service
	info           *info.Service
	analog         *analog.Service
	messenger      *messenger.Service
	traceEventFunc func(event TraceInfoEvent)
	TsUser         string
	TsPassword     string
}

// NewClient creates a new instance of Tradesoft API client
func NewClient(baseURL string, user, password string) *Client {
	httpClient := resty.New()

	client := Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		TsUser:     user,
		TsPassword: password,
		provider:   provider.NewService(httpClient, baseURL, user, password),
		info:       info.NewService(httpClient, baseURL, user, password),
		analog:     analog.NewService(httpClient, baseURL, user, password),
		messenger:  messenger.NewService(httpClient, baseURL, user, password),
	}
	client.httpClient.SetBasicAuth(user, password)
	return &client
}

func (c *Client) TraceOff() {
	c.httpClient.DisableTrace()
}

func (c *Client) TraceOn(event func(event TraceInfoEvent)) {
	c.traceEventFunc = event
	c.httpClient.OnSuccess(func(client *resty.Client, response *resty.Response) {
		c.traceEventFunc(TraceInfoEvent{
			TraceInfo:  response.Request.TraceInfo(),
			Url:        response.Request.URL,
			Method:     response.Request.Method,
			BodySize:   response.Size(),
			StatusCode: response.StatusCode(),
		})
	})
	c.httpClient.EnableTrace()
}

func NewClientDefault(user, password string) *Client {
	return NewClient(TradeSoftUrl, user, password)
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
