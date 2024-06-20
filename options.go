package hbookerLib

import (
	"regexp"
)

type Options interface {
	Apply(client *Client)
}
type OptionFunc func(client *Client)

func (f OptionFunc) Apply(client *Client) {
	f(client)
}

func WithLoginToken(loginToken string) Options {
	return OptionFunc(func(client *Client) {
		client.Authenticate.SetLoginToken(loginToken)
	})
}
func WithAccount(account string) Options {
	return OptionFunc(func(client *Client) {
		client.Authenticate.SetAccount(account)
	})
}
func WithVersion(version string) Options {
	return OptionFunc(func(client *Client) {
		client.Authenticate.SetAppVersion(version)
	})
}
func WithIos() Options {
	return OptionFunc(func(client *Client) {
		client.ios = true
	})

}
func WithRetryCount(retryCount int) Options {
	return OptionFunc(func(client *Client) {
		if retryCount > 0 {
			client.retryCount = retryCount
		}
	})
}
func WithDebug() Options {
	return OptionFunc(func(client *Client) {
		client.debug = !client.debug
	})
}

func WithOutputDebug() Options {
	return OptionFunc(func(client *Client) {
		client.outputDebug = !client.outputDebug
	})
}

func WithProxyURL(proxyURL string) Options {
	return OptionFunc(func(client *Client) {
		if regexp.MustCompile(`^http[s]?://[a-zA-Z0-9.-]+(:\d+)?$`).MatchString(proxyURL) {
			client.proxyURL = proxyURL
		}
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		if regexp.MustCompile(`^https?://[a-zA-Z0-9.-]+(:\d+)?(/.*)?$`).MatchString(apiBaseURL) {
			client.baseURL = apiBaseURL
		}
	})
}

func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.apiKey = androidApiKey
	})
}
func WithDeviceToken(deviceToken string) Options {
	return OptionFunc(func(client *Client) {
		client.Authenticate.SetDeviceToken(deviceToken)
	})
}
