package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/imroc/req/v3"
)

type Client struct {
	version       string
	baseURL       string
	androidApiKey string
	deviceToken   string
	LoginToken    string
	Account       string
	debug         bool
	retryCount    int
	outputDebug   bool
	proxyURL      string
	HttpsClient   *req.Client
}

type API struct {
	HttpRequest *req.Request
}

func defaultConfig() *Client {
	client := &Client{HttpsClient: req.NewClient()}
	for _, option := range []Options{
		WithVersion(version),
		WithRetryCount(retryCount),
		WithDeviceToken(deviceToken),
		WithAndroidApiKey(androidApiKey),
		WithAPIBaseURL(urlconstants.WEB_SITE),
	} {
		option.Apply(client)
	}
	return client
}

func NewClient(options ...Options) *Client {
	client := defaultConfig()
	for _, option := range options {
		option.Apply(client)
	}

	return client
}
