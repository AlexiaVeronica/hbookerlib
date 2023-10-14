package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookerapi"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"sync"
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
func (client *Client) Download(bookId string,
	continueFunc func(hbookermodel.ChapterList) bool,
	contentFunc func(hbookermodel.ChapterList, string),
) error {

	divisionList, err := client.API.GetDivisionListByBookId(bookId)
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	ch := make(chan int, 42)
	for _, division := range divisionList {
		var count int
		for _, chapter := range division.ChapterList {
			ch <- 1
			wg.Add(1)
			go func(chapter hbookermodel.ChapterList, wg *sync.WaitGroup, ch chan int) {
				defer func() {
					lock := sync.Mutex{}
					lock.Lock()
					count++
					fmt.Printf("downloaded: %d/%d\r", count, len(division.ChapterList))
					lock.Unlock()
					<-ch
					wg.Done()
				}()

				if !continueFunc(chapter) {
					return
				}
				content, ok := client.NewGetContent(chapter.ChapterID)
				if ok != nil {
					fmt.Println(ok)
				} else {
					contentFunc(chapter, content)
				}
			}(chapter, &wg, ch)
		}
	}
	wg.Wait()
	return nil
}
