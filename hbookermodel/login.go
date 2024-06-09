package hbookermodel

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Login struct {
	Tip
	Data LoginData `json:"data"`
}

type LoginData struct {
	LoginToken string     `json:"login_token"`
	UserCode   string     `json:"user_code"`
	ReaderInfo ReaderInfo `json:"reader_info"`
	PropInfo   PropInfo   `json:"prop_info"`
	IsSetYoung string     `json:"is_set_young"`
}

type Register struct {
	Tip
	Data LoginData `json:"data"`
}
type Authenticate struct {
	AppVersion  string `json:"app_version"`
	DeviceToken string `json:"device_token"`
	LoginToken  string `json:"login_token"`
	Account     string `json:"account"`
}

func (authenticate *Authenticate) SetAppVersion(appVersion string) *Authenticate {
	// Regular expression to match semantic versioning (e.g., 1.0.0, 2.9.290)
	if regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(appVersion) {
		authenticate.AppVersion = appVersion
	}
	return authenticate
}
func (authenticate *Authenticate) SetDeviceToken(deviceToken string) *Authenticate {
	authenticate.DeviceToken = deviceToken
	return authenticate
}
func (authenticate *Authenticate) SetLoginToken(loginToken string) *Authenticate {
	if len(loginToken) == 32 {
		authenticate.LoginToken = loginToken
	}
	return authenticate
}
func (authenticate *Authenticate) SetAccount(account string) *Authenticate {
	if unquoted, err := strconv.Unquote(fmt.Sprintf(`"%s"`, account)); err == nil {
		account = unquoted
	}
	// Check if the (possibly decoded) string contains "书客".
	if strings.Contains(account, "书客") {
		authenticate.Account = account
	}
	return authenticate
}
func (authenticate *Authenticate) GetAuthenticate() map[string]string {
	var result map[string]string
	m, _ := json.Marshal(authenticate)
	err := json.Unmarshal(m, &result)
	if err != nil {
		log.Panicln("Error unmarshalling Authenticate:", err)
		return nil
	}
	return result
}
