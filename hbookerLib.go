package hbookerLib

import (
	"github.com/imroc/req/v3"
)

const (
	Version       = "2.9.290"
	DeviceToken   = "ciweimao_"
	AndroidApiKey = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	APIBaseURL    = "https://app.hbooker.com/"
	UserAgent     = "Android com.kuangxiangciweimao.novel "
)

type Client struct {
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
	HttpsClient   *req.Client
}

type API struct {
	HttpRequest *req.Request
}

func defaultConfig() *Client {
	client := &Client{HttpsClient: req.NewClient()}
	for _, option := range []Options{
		WithVersion(Version),
		WithDeviceToken(DeviceToken),
		WithAndroidApiKey(AndroidApiKey),
		WithAPIBaseURL(APIBaseURL),
		WithUserAgent(UserAgent),
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
	httpRequest := client.HttpsClient.
		SetCommonRetryCount(5).
		SetBaseURL(client.APIBaseURL).SetResponseBodyTransformer(func(rawBody []byte, _ *req.Request, _ *req.Response) ([]byte, error) {
		return aesDecrypt(string(rawBody), client.AndroidApiKey)
	}).R()

	httpRequest.SetFormData(map[string]string{
		"app_version":  client.Version,
		"device_token": client.DeviceToken,
		"login_token":  client.LoginToken,
		"account":      client.Account,
	})
	httpRequest.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   client.UserAgent + client.Version,
	})

	return &API{HttpRequest: httpRequest}
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
