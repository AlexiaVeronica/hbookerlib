package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/imroc/req/v3"
)

const (
	version       = "2.9.290"
	deviceToken   = "ciweimao_"
	androidApiKey = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	retryCount    = 5
	userAgent     = "Android com.kuangxiangciweimao.novel "
)

type Client struct {
	version       string
	baseURL       string
	androidApiKey string
	deviceToken   string
	LoginToken    string
	Account       string
	debug         bool
	retryCount    int
	outputDebug   bool
	proxyURL      string
	HttpsClient   *req.Client
}

type API struct {
	HttpRequest *req.Request
}

func defaultConfig() *Client {

	client := &Client{HttpsClient: req.NewClient()}
	for _, option := range []Options{
		WithVersion(version),
		WithRetryCount(retryCount),
		WithDeviceToken(deviceToken),
		WithAndroidApiKey(androidApiKey),
		WithAPIBaseURL(urlconstants.WEB_SITE),
	} {
		option.Apply(client)
	}
	return client
}
func NewClient(options ...Options) *Client {
	client := defaultConfig()
	for _, option := range options {
		option.Apply(client)
	}

	return client
}

func (client *Client) API() *API {
	if client.debug {
		client.HttpsClient.DevMode()
	}
	if client.outputDebug {
		client.HttpsClient.EnableDumpAllToFile("hbookerLib_debug.log")
	}
	if client.proxyURL != "" {
		client.HttpsClient.SetProxyURL(client.proxyURL)
	}
	httpRequest := client.HttpsClient.
		SetCommonRetryCount(client.retryCount).
		SetBaseURL(client.baseURL).SetResponseBodyTransformer(func(rawBody []byte, _ *req.Request, _ *req.Response) ([]byte, error) {
		return aesDecrypt(string(rawBody), client.androidApiKey)
	}).R()

	httpRequest.SetFormData(map[string]string{
		"app_version":  client.version,
		"device_token": client.deviceToken,
		"login_token":  client.LoginToken,
		"account":      client.Account,
	})
	httpRequest.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   userAgent + client.version,
	})
	return &API{HttpRequest: httpRequest}
}
func (api *API) DeleteValue(deleteValue string) *API {
	if api.HttpRequest.FormData != nil {
		delete(api.HttpRequest.FormData, deleteValue)
	}
	return api
}

//
//func (client *Client) SetDefaultParams(account, loginToken string) {
//	client.API.HttpClient.Account = account
//	client.API.HttpClient.LoginToken = loginToken
//}

//func (client *Client) NewGetContent(chapterId string) (string, error) {
//	key, err := client.API.GetChapterKey(chapterId)
//	if err != nil {
//		return "", err
//	}
//	content, err := client.API.GetChapterContentAPI(chapterId, key)
//	if err != nil {
//		return "", err
//	}
//	return string(hbookerapi.HbookerDecode(content.TxtContent, key)), nil
//
//}
//func (client *Client) Download(bookId string,
//	continueFunc func(hbookermodel.ChapterList) bool,
//	contentFunc func(hbookermodel.ChapterList, string),
//) error {
//
//	divisionList, err := client.API.GetDivisionListByBookId(bookId)
//	if err != nil {
//		return err
//	}
//	wg := sync.WaitGroup{}
//	ch := make(chan int, 42)
//	for _, division := range divisionList {
//		var count int
//		for _, chapter := range division.ChapterList {
//			ch <- 1
//			wg.Add(1)
//			go func(chapter hbookermodel.ChapterList, wg *sync.WaitGroup, ch chan int) {
//				defer func() {
//					lock := sync.Mutex{}
//					lock.Lock()
//					count++
//					fmt.Printf("downloaded: %d/%d\r", count, len(division.ChapterList))
//					lock.Unlock()
//					<-ch
//					wg.Done()
//				}()
//
//				if !continueFunc(chapter) {
//					return
//				}
//				content, ok := client.NewGetContent(chapter.ChapterID)
//				if ok != nil {
//					fmt.Println(ok)
//				} else {
//					contentFunc(chapter, content)
//				}
//			}(chapter, &wg, ch)
//		}
//	}
//	wg.Wait()
//	return nil
//}
