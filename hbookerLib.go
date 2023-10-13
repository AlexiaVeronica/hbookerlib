package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookerapi"
)

type Client struct {
	API *hbookerapi.API
}

func NewClient(options ...Options) *Client {
	client := &Client{
		API: &hbookerapi.API{
			HttpClient: hbookerapi.HttpsClient{
				Debug:         false,
				OutputDebug:   false,
				Version:       "2.9.290",
				DeviceToken:   "ciweimao_",
				AndroidApiKey: "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn",
				APIBaseURL:    "https://app.hbooker.com/",
				UserAgent:     "Android com.kuangxiangciweimao.novel ",
			},
		},
	}
	for _, option := range options {
		option.Apply(client)
	}
	return client
}

func (client *Client) SetDefaultParams(account, loginToken string) {
	client.API.HttpClient.Account = account
	client.API.HttpClient.LoginToken = loginToken
}
func (client *Client) NewGetContent(chapterId string) (string, error) {
	key, err := client.API.GetChapterKey(chapterId)
	if err != nil {
		return "", err
	}
	content, err := client.API.GetChapterContentAPI(chapterId, key)
	if err != nil {
		return "", err
	}
	return string(hbookerapi.HbookerDecode(content.TxtContent, key)), nil

}
