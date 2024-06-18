package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/AlexiaVeronica/req/v3"
	"github.com/google/uuid"
)

type Client struct {
	baseURL       string
	androidApiKey string
	debug         bool
	retryCount    int
	outputDebug   bool
	proxyURL      string
	HttpsClient   *req.Client
	Authenticate  *hbookermodel.Authenticate
}

type API struct {
	HttpRequest *req.Request
}

func defaultConfig() *Client {
	return &Client{
		HttpsClient: req.NewClient(),
		Authenticate: &hbookermodel.Authenticate{
			AppVersion:  version,
			DeviceToken: deviceToken,
		},
		retryCount:    retryCount,
		androidApiKey: androidApiKey,
		baseURL:       urlconstants.WEB_SITE,
	}
}
func defaultIosConfig() *Client {
	return &Client{
		HttpsClient: req.NewClient(),
		Authenticate: &hbookermodel.Authenticate{
			AppVersion:  versionIos,
			DeviceToken: deviceIosToken + uuid.New().String(),
		},
		retryCount:    retryCount,
		androidApiKey: androidApiKey,
		baseURL:       urlconstants.WEB_SITE,
	}
}
func NewClient(options ...Options) *Client {
	client := defaultIosConfig()
	for _, option := range options {
		option.Apply(client)
	}
	return client
}
func (client *Client) SetToken(account, loginToken string) *Client {
	WithLoginToken(loginToken).Apply(client)
	WithAccount(account).Apply(client)
	return client
}
