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
			client.API.HttpClient.LoginToken = loginToken
		}
	})
}
func WithAccount(account string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.Account = account
	})
}
func WithAccountAndLoginToken(account, loginToken string) Options {
	return OptionFunc(func(client *Client) {
		WithAccount(account).Apply(client)
		WithLoginToken(loginToken).Apply(client)
		_, err := client.API.GetUserInfo()
		if err != nil {
			log.Println(err)
			client.API.HttpClient.Account = ""
			client.API.HttpClient.LoginToken = ""
			return
		}

	})
}
func WithVersion(version string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.Version = version
	})
}
func WithDebug() Options {
	return OptionFunc(func(client *Client) {
		if client.API.HttpClient.Debug {
			client.API.HttpClient.Debug = false
		} else {
			client.API.HttpClient.Debug = true
		}
	})
}

func WithOutputDebug() Options {
	return OptionFunc(func(client *Client) {
		if client.API.HttpClient.OutputDebug {
			client.API.HttpClient.OutputDebug = false
		} else {
			client.API.HttpClient.OutputDebug = true
		}
	})
}
func WithProxyURLArray(proxyURLArray []string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.ProxyURLArray = proxyURLArray
	})
}
func WithProxyURL(proxyURL string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.ProxyURL = proxyURL
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.APIBaseURL = apiBaseURL
	})
}
func WithUserAgent(userAgent string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.UserAgent = userAgent
	})
}
func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.AndroidApiKey = androidApiKey
	})
}
