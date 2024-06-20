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
	return &Client{
		HttpsClient: defaultReqClient(),
		Authenticate: &hbookermodel.Authenticate{
			AppVersion:  versionAndroid,
			DeviceToken: deviceToken,
		},
		retryCount: retryCount,
		apiKey:     apiKey,
		baseURL:    urlconstants.WEB_SITE,
	}
}
func defaultIosConfig() *Client {
	return &Client{
		HttpsClient: defaultReqClient(),
		Authenticate: &hbookermodel.Authenticate{
			AppVersion:  versionIos,
			DeviceToken: deviceIosToken + uuid.New().String(),
		},
		retryCount: retryCount,
		apiKey:     apiKey,
		baseURL:    urlconstants.WEB_SITE,
	}
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
