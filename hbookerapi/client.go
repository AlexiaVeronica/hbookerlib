package hbookerapi

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
)

type HttpsClient struct {
	Version       string
	APIBaseURL    string
	UserAgent     string
	AndroidApiKey string
	DeviceToken   string
	LoginToken    string
	Account       string
	Debug         bool
	OutputDebug   bool
	ProxyURL      string
	ProxyURLArray []string
}

func (httpsClient *HttpsClient) defaultHeader() map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Android com.kuangxiangciweimao.novel " + httpsClient.Version,
	}
}
func (httpsClient *HttpsClient) defaultFormData() map[string]string {
	return map[string]string{"app_version": httpsClient.Version, "device_token": httpsClient.DeviceToken, "login_token": httpsClient.LoginToken, "account": httpsClient.Account}
}
func (httpsClient *HttpsClient) NewDefault() *req.Client {
	c := req.C().
		SetCommonRetryCount(5).
		SetTimeout(30 * time.Second).
		SetBaseURL(httpsClient.APIBaseURL).
		SetCommonHeaders(httpsClient.defaultHeader())
	if httpsClient.ProxyURL != "" {
		c.SetProxyURL(httpsClient.ProxyURL)
	}
	if len(httpsClient.ProxyURLArray) > 0 {
		c.SetProxyURL(httpsClient.ProxyURLArray[time.Now().UnixNano()%int64(len(httpsClient.ProxyURLArray))])
	}
	if httpsClient.Debug {
		c.DevMode()
		if httpsClient.OutputDebug {
			c.EnableDumpAllToFile("hbooker.log")
		}
	}
	return c
}
func (httpsClient *HttpsClient) Post(path string, params map[string]string, model any) (*req.Response, error) {
	if params == nil {
		params = httpsClient.defaultFormData()
	} else {
		for k, v := range httpsClient.defaultFormData() {
			params[k] = v
		}
	}
	response, err := httpsClient.NewDefault().R().SetFormData(params).Post(path)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccessState() {
		return nil, fmt.Errorf("response is not success state: %v", response.String())
	}
	if model != nil {
		result := response.String()
		if match, _ := regexp.MatchString(`^\{.*}$`, result); !match {
			result = string(HbookerDecode(response.String(), httpsClient.AndroidApiKey))
		}
		err = json.Unmarshal([]byte(result), model)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal error: %v", err)
		}
	}
	return response, nil
}
func (httpsClient *HttpsClient) Get(path string, params map[string]string, model any) (*req.Response, error) {
	newDefault := httpsClient.NewDefault().R().SetQueryParams(params)
	response, err := newDefault.Get(path)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccessState() {
		return nil, fmt.Errorf("response is not success state: %v", response.String())
	}
	if model != nil {
		result := response.String()
		if match, _ := regexp.MatchString(`^\{.*}$`, result); !match {
			result = string(HbookerDecode(response.String(), httpsClient.AndroidApiKey))
		}
		err = json.Unmarshal([]byte(result), model)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal error: %v", err)
		}
	}
	return response, nil
}
