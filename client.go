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
func defaultAndroidConfig() *Client {
	client := &Client{HttpsClient: defaultReqClient()}
	options := []Options{
		WithVersion(versionAndroid),
		WithDeviceToken(deviceToken),
		WithRetryCount(retryCount),
		WithApiKey(apiKey),
		WithAPIBaseURL(urlconstants.WEB_SITE),
	}
	for _, option := range options {
		option.Apply(client)
	}
	return client
}
func defaultIosConfig() *Client {
	client := &Client{HttpsClient: defaultReqClient()}
	options := []Options{
		WithVersion(versionIos),
		WithDeviceToken(deviceIosToken + uuid.New().String()),
		WithRetryCount(retryCount),
		WithApiKey(apiKey),
		WithAPIBaseURL(urlconstants.WEB_SITE),
	}
	for _, option := range options {
		option.Apply(client)
	}
	return client
}
func NewClient(options ...Options) *Client {
	if len(options) == 0 {
		return defaultIosConfig()
	} else {
		client := defaultIosConfig()
		for _, option := range options {
			option.Apply(client)
		}
		if !client.ios {
			client = defaultAndroidConfig()
			for _, option := range options {
				option.Apply(client)
			}
		}
		return client
	}

}
func (client *Client) SetToken(account, loginToken string) *Client {
	WithLoginToken(loginToken).Apply(client)
	WithAccount(account).Apply(client)
	return client
}
