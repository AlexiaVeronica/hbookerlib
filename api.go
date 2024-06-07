package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/imroc/req/v3"
	"strconv"
	"time"
)

// 泛型函数，用于处理响应的解码和错误检查
func handleResponse[T any](res *req.Response, err error, data *T) (*T, error) {
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, fmt.Errorf("response is nil")
	}
	if err = res.UnmarshalJson(data); err != nil {
		return nil, err
	}

	// 使用类型断言访问 Code 和 Tip 字段
	v, ok := any(data).(interface {
		GetCode() string
		GetTip() string
	})
	if !ok {
		return nil, fmt.Errorf("response does not implement required methods")
	}

	if v.GetCode() != "100000" {
		return nil, fmt.Errorf("error: %s", v.GetTip())
	}

	return data, nil
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
	httpRequest.SetHeaders(map[string]string{"Content-Type": postContentType, "User-Agent": userAgent + client.version})

	return &API{HttpRequest: httpRequest}
}

func (api *API) DeleteValue(deleteValue string) *API {
	if api.HttpRequest.FormData != nil {
		delete(api.HttpRequest.FormData, deleteValue)
	}
	return api
}

func (api *API) GetBookInfo(bookId string) (*hbookermodel.Detail, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"book_id": bookId}).Post(urlconstants.BookGetInfoById)
	return handleResponse(res, err, &hbookermodel.Detail{})
}

func (api *API) GetUserInfo() (*hbookermodel.UserInfo, error) {
	res, err := api.HttpRequest.Post(urlconstants.MY_DETAILS_INFO)
	return handleResponse(res, err, &hbookermodel.UserInfo{})
}

func (api *API) GetDivisionListByBookId(bookId string) (*hbookermodel.NewDivisionModel, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"book_id": bookId}).Post(urlconstants.GetDivisionListNew)
	return handleResponse(res, err, &hbookermodel.NewDivisionModel{})
}

func (api *API) GetChapterKey(chapterId string) (*hbookermodel.ContentKey, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId}).Post(urlconstants.GetChapterKey)
	return handleResponse(res, err, &hbookermodel.ContentKey{})
}

func (api *API) GetChapterContentAPI(chapterId, chapterKey string) (*hbookermodel.ChapterInfo, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId, "chapter_command": chapterKey}).Post(urlconstants.GetCptIfm)
	content, ok := handleResponse(res, err, &hbookermodel.Content{})
	if ok != nil {
		return nil, ok
	}
	contentRaw, err := aesDecrypt(content.Data.ChapterInfo.TxtContent, chapterKey)
	if err != nil {
		return nil, err
	}
	content.Data.ChapterInfo.TxtContent = string(contentRaw)
	return &content.Data.ChapterInfo, nil

}

// Deprecated: MySignLogin is deprecated, hbooker has joined login verification, so this method is no longer available
func (api *API) MySignLogin(username, password, validate, challenge string) (*hbookermodel.Login, error) {
	params := map[string]string{"login_name": username, "passwd": password}
	if validate != "" {
		params["geetest_seccode"] = validate + "|jordan"
		params["geetest_validate"] = validate
		params["geetest_challenge"] = challenge
	}
	res, err := api.DeleteValue("login_token").DeleteValue("account").
		HttpRequest.SetFormData(params).Post(urlconstants.MySignLogin)
	return handleResponse(res, err, &hbookermodel.Login{})
}

func (api *API) GetBuyChapterAPI(chapterId string) (*hbookermodel.ContentBuy, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId, "shelf_id": ""}).Post(urlconstants.ChapterBuy)
	return handleResponse(res, err, &hbookermodel.ContentBuy{})
}

func (api *API) GetAutoSignAPI(device string) (*hbookermodel.Register, error) {
	if !checkDeviceRegex.MatchString(device) {
		return nil, fmt.Errorf("get auto sign error: %s", "device is not valid")
	}
	res, err := api.DeleteValue("login_token").DeleteValue("account").HttpRequest.
		SetFormData(map[string]string{"uuid": "android" + device, "gender": "1", "channel": "PCdownloadC"}).Post(urlconstants.SIGNUP)
	return handleResponse(res, err, &hbookermodel.Register{})
}

func (api *API) GetUseGeetestAPI(loginName string) (*hbookermodel.Geetest, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"login_name": loginName}).Post(urlconstants.UseGeetest)
	return handleResponse(res, err, &hbookermodel.Geetest{})
}

func (api *API) GetGeetestRegisterAPI(UserID string) (*hbookermodel.GeetestFirstRegisterStruct, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"user_id": UserID, "t": strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}).Post(urlconstants.GeetestRegister)
	return handleResponse(res, err, &hbookermodel.GeetestFirstRegisterStruct{})
}

func (api *API) GetBookcaseAPI(shelfId string) (*hbookermodel.Bookcase, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"shelf_id": shelfId, "direction": "prev", "last_mod_time": "0"}).Post(urlconstants.BookshelfGetShelfBookList)
	return handleResponse(res, err, &hbookermodel.Bookcase{})
}

func (api *API) GetBookShelfInfoAPI() (*hbookermodel.Bookshelf, error) {
	res, err := api.HttpRequest.Post(urlconstants.BookshelfGetShelfList)
	return handleResponse(res, err, &hbookermodel.Bookshelf{})
}

func (api *API) GetSearchBooksAPI(keyword string, page any) (*hbookermodel.Search, error) {
	res, err := api.HttpRequest.SetFormData(map[string]string{"count": "10", "page": fmt.Sprintf("%v", page), "category_index": "0", "key": keyword}).
		Post(urlconstants.BookcityGetFilterList)
	return handleResponse(res, err, &hbookermodel.Search{})
}
