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
		client.version = version
	})
}
func WithRetryCount(retryCount int) Options {
	return OptionFunc(func(client *Client) {
		client.retryCount = retryCount
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
		client.proxyURL = proxyURL
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		client.baseURL = apiBaseURL
	})
}

func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.androidApiKey = androidApiKey
	})
}
func WithDeviceToken(deviceToken string) Options {
	return OptionFunc(func(client *Client) {
		client.deviceToken = deviceToken
	})
}
