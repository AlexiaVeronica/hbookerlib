package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/AlexiaVeronica/req/v3"
	"github.com/google/uuid"
)

type Client struct {
	ios          bool
	baseURL      string
	apiKey       string
	debug        bool
	retryCount   int
	outputDebug  bool
	proxyURL     string
	HttpsClient  *req.Client
	Authenticate *hbookermodel.Authenticate
}

type API struct {
	HttpRequest *req.Request
}

func defaultReqClient() *req.Client {
	return req.NewClient().SetCommonHeader("Content-Type", postContentType)
}

func configureClient(client *Client, options []Options) {
	for _, option := range options {
		option.Apply(client)
	}
}

func defaultConfig(version, deviceToken string) *Client {
	client := &Client{HttpsClient: defaultReqClient()}
	options := []Options{
		WithVersion(version),
		WithDeviceToken(deviceToken),
		WithRetryCount(retryCount),
		WithApiKey(apiKey),
		WithAPIBaseURL(urlconstants.WEB_SITE),
	}
	configureClient(client, options)
	return client
}

func defaultAndroidConfig() *Client {
	return defaultConfig(versionAndroid, deviceToken)
}

func defaultIosConfig() *Client {
	return defaultConfig(versionIos, deviceIosToken+uuid.New().String())
}

func NewClient(options ...Options) *Client {
	client := defaultIosConfig()
	if len(options) > 0 {
		configureClient(client, options)
		if !client.ios {
			client = defaultAndroidConfig()
			configureClient(client, options)
		}
	}
	return client
}

func (client *Client) SetLoginToken(account, loginToken string) *Client {
	WithLoginToken(loginToken).Apply(client)
	WithAccount(account).Apply(client)
	return client
}
