package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/imroc/req/v3"
	"strconv"
	"time"
)

type Request[T any] struct {
	HttpRequest *req.Request
}

func (request *Request[T]) handleResponse(url string, formData map[string]string) (*T, error) {
	res, err := request.HttpRequest.SetFormData(formData).Post(url)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}

	data := new(T)
	if err = res.UnmarshalJson(data); err != nil {
		return nil, err
	}

	if response, ok := any(data).(interface {
		GetCode() string
		GetTip() string
		IsSuccess() bool
	}); ok && !response.IsSuccess() {
		return nil, fmt.Errorf("error: %s", response.GetTip())
	} else if !ok {
		return nil, fmt.Errorf("response does not implement required methods")
	}

	return data, nil
}

func newRequest[T any](HttpRequest *req.Request) *Request[T] {
	return &Request[T]{HttpRequest: HttpRequest}
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
		SetBaseURL(client.baseURL).
		SetResponseBodyTransformer(func(rawBody []byte, _ *req.Request, _ *req.Response) ([]byte, error) {
			return aesDecrypt(string(rawBody), client.androidApiKey)
		}).R()

	httpRequest.SetFormData(map[string]string{
		"app_version":  client.version,
		"device_token": client.deviceToken,
		"login_token":  client.LoginToken,
		"account":      client.Account,
	}).SetHeaders(map[string]string{
		"Content-Type": postContentType,
		"User-Agent":   userAgent + client.version,
	})

	return &API{HttpRequest: httpRequest}
}

func (api *API) GetBookInfo(bookId string) (*hbookermodel.Detail, error) {
	data := map[string]string{"book_id": bookId}
	return newRequest[hbookermodel.Detail](api.HttpRequest).handleResponse(urlconstants.GetInfoById, data)
}

func (api *API) DeleteValue(deleteValue string) *API {
	if api.HttpRequest.FormData != nil {
		delete(api.HttpRequest.FormData, deleteValue)
	}
	return api
}

func (api *API) GetUserInfo() (*hbookermodel.UserInfo, error) {
	return newRequest[hbookermodel.UserInfo](api.HttpRequest).handleResponse(urlconstants.MyDetailsInfo, nil)
}

func (api *API) GetDivisionListByBookId(bookId string) (*hbookermodel.Division, error) {
	data := map[string]string{"book_id": bookId}
	return newRequest[hbookermodel.Division](api.HttpRequest).handleResponse(urlconstants.GetUpdatedChapterByDivisionNew, data)
}

func (api *API) GetChapterCmd(chapterId string) (*hbookermodel.ChapterCmd, error) {
	data := map[string]string{"chapter_id": chapterId}
	return newRequest[hbookermodel.ChapterCmd](api.HttpRequest).handleResponse(urlconstants.GetChapterCmd, data)
}

func (api *API) GetChapterContentAPI(chapterId, chapterKey string) (*hbookermodel.ChapterInfo, error) {
	data := map[string]string{"chapter_id": chapterId, "chapter_command": chapterKey}
	content, err := newRequest[hbookermodel.Content](api.HttpRequest).handleResponse(urlconstants.GetCptIfm, data)
	if err != nil {
		return nil, err
	}

	contentRaw, err := aesDecrypt(content.Data.ChapterInfo.TxtContent, chapterKey)
	if err != nil {
		return nil, fmt.Errorf("aesDecrypt error: %v", err)
	}
	content.Data.ChapterInfo.TxtContent = string(contentRaw)
	return &content.Data.ChapterInfo, nil
}

// Deprecated: MySignLogin is deprecated, hbooker has joined login verification, so this method is no longer available
func (api *API) MySignLogin(username, password, validate, challenge string) (*hbookermodel.Login, error) {
	data := map[string]string{"login_name": username, "passwd": password}
	if validate != "" {
		data["geetest_seccode"] = validate + "|jordan"
		data["geetest_validate"] = validate
		data["geetest_challenge"] = challenge
	}
	return newRequest[hbookermodel.Login](api.HttpRequest).handleResponse(urlconstants.MySignLogin, data)
}

func (api *API) GetBuyChapterAPI(chapterId string) (*hbookermodel.ContentBuy, error) {
	data := map[string]string{"chapter_id": chapterId, "shelf_id": ""}
	return newRequest[hbookermodel.ContentBuy](api.HttpRequest).handleResponse(urlconstants.ChapterBuy, data)
}

func (api *API) GetAutoSignAPI(device string) (*hbookermodel.Register, error) {
	data := map[string]string{"uuid": "android" + device, "gender": "1", "channel": "PCdownloadC"}
	return newRequest[hbookermodel.Register](api.HttpRequest).handleResponse(urlconstants.SIGNUP, data)
}

func (api *API) GetUseGeetestAPI(loginName string) (*hbookermodel.Geetest, error) {
	data := map[string]string{"login_name": loginName}
	return newRequest[hbookermodel.Geetest](api.HttpRequest).handleResponse(urlconstants.UseGeetest, data)
}

func (api *API) GetGeetestRegisterAPI(UserID string) (*hbookermodel.GeetestFirstRegisterStruct, error) {
	data := map[string]string{"user_id": UserID, "t": strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}
	return newRequest[hbookermodel.GeetestFirstRegisterStruct](api.HttpRequest).handleResponse(urlconstants.GeetestRegister, data)
}

func (api *API) GetBookcaseAPI(shelfId string) (*hbookermodel.Bookcase, error) {
	data := map[string]string{"shelf_id": shelfId, "direction": "prev", "last_mod_time": "0"}
	return newRequest[hbookermodel.Bookcase](api.HttpRequest).handleResponse(urlconstants.BookshelfGetShelfBookList, data)
}

func (api *API) GetBookShelfInfoAPI() (*hbookermodel.Bookshelf, error) {
	return newRequest[hbookermodel.Bookshelf](api.HttpRequest).handleResponse(urlconstants.BookshelfGetShelfList, nil)
}

func (api *API) GetSearchBooksAPI(keyword string, page any) (*hbookermodel.Search, error) {
	data := map[string]string{"count": "10", "page": fmt.Sprintf("%v", page), "category_index": "0", "key": keyword}
	return newRequest[hbookermodel.Search](api.HttpRequest).handleResponse(urlconstants.GetFilterSearchBookList, data)
}
