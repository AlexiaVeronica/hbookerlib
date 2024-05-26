package hbookerLib

import (
	"log"
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
		if len(loginToken) != 32 {
			log.Println("LoginToken is must be 32 length, please check it.")
		} else {
			client.LoginToken = loginToken
		}
	})
}
func WithAccount(account string) Options {
	return OptionFunc(func(client *Client) {
		client.Account = account
	})
}
func WithVersion(version string) Options {
	return OptionFunc(func(client *Client) {
		client.Version = version
	})
}
func WithDebug() Options {
	return OptionFunc(func(client *Client) {
		client.Debug = !client.Debug
	})
}

func WithOutputDebug() Options {
	return OptionFunc(func(client *Client) {
		client.OutputDebug = !client.OutputDebug
	})
}
func WithProxyURLArray(proxyURLArray []string) Options {
	return OptionFunc(func(client *Client) {
		client.ProxyURLArray = proxyURLArray
	})
}
func WithProxyURL(proxyURL string) Options {
	return OptionFunc(func(client *Client) {
		client.ProxyURL = proxyURL
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		client.APIBaseURL = apiBaseURL
	})
}
func WithUserAgent(userAgent string) Options {
	return OptionFunc(func(client *Client) {
		client.UserAgent = userAgent
	})
}
func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.AndroidApiKey = androidApiKey
	})
}
func WithDeviceToken(deviceToken string) Options {
	return OptionFunc(func(client *Client) {
		client.DeviceToken = deviceToken
	})
}
