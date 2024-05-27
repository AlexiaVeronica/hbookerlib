package hbookerLib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
		if len(loginToken) == 32 {
			client.LoginToken = loginToken
		}
	})
}
func WithAccount(account string) Options {
	return OptionFunc(func(client *Client) {
		if unquoted, err := strconv.Unquote(fmt.Sprintf(`"%s"`, account)); err == nil {
			account = unquoted
		}
		// Check if the (possibly decoded) string contains "书客".
		if strings.Contains(account, "书客") {
			client.Account = account
		}
	})
}
func WithVersion(version string) Options {
	return OptionFunc(func(client *Client) {
		// Regular expression to match semantic versioning (e.g., 1.0.0, 2.9.290)
		if regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(version) {
			client.version = version
		}
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
		client.androidApiKey = androidApiKey
	})
}
func WithDeviceToken(deviceToken string) Options {
	return OptionFunc(func(client *Client) {
		client.deviceToken = deviceToken
	})
}
