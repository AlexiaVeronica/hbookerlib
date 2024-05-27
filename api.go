package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/hbookerLib/urlconstants"
	"github.com/imroc/req/v3"
	"strconv"
	"strings"
	"time"
)

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
func (api *API) GetBookInfo(bookId string) (*hbookermodel.BookInfo, error) {
	var book hbookermodel.Detail
	res, err := api.HttpRequest.SetFormData(map[string]string{"book_id": bookId}).Post(urlconstants.BookGetInfoById)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&book)
	if book.Code != "100000" {
		return nil, fmt.Errorf("get book information error: %s", book.Tip)
	}
	if book.Data.BookInfo.BookName == "" {
		return nil, fmt.Errorf("get book information error: %s", "book name is empty")
	}
	var tagList []string
	for _, tag := range book.Data.BookInfo.TagList {
		tagList = append(tagList, tag.TagName)
	}
	if len(tagList) > 0 {
		book.Data.BookInfo.Tag = strings.Join(tagList, ",")
	}
	return &book.Data.BookInfo, nil
}

func (api *API) GetUserInfo() (*hbookermodel.UserInfoData, error) {
	var m hbookermodel.UserInfo
	res, err := api.HttpRequest.Post(urlconstants.MY_DETAILS_INFO)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&m)
	if m.Code != "100000" {
		return nil, fmt.Errorf("get user info error: %s", m.Tip)
	}
	return &m.Data, nil
}

func (api *API) GetDivisionListByBookId(bookId string) ([]hbookermodel.VolumeList, error) {
	var divisionList hbookermodel.NewVolumeList
	res, err := api.HttpRequest.SetFormData(map[string]string{"book_id": bookId}).Post(urlconstants.GetDivisionListNew)

	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&divisionList)
	if divisionList.Code != "100000" {
		return nil, fmt.Errorf("get division list error: %s", divisionList.Tip)
	}
	if len(divisionList.Data.ChapterList) == 0 {
		return nil, fmt.Errorf("get division list error: %s", "division list is empty")
	}
	return divisionList.Data.ChapterList, nil
}

func (api *API) GetChapterKey(chapterId string) (string, error) {
	var m hbookermodel.ContentKey
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId}).Post(urlconstants.GetChapterKey)
	if err != nil {
		return "", err
	}
	res.UnmarshalJson(&m)
	if m.Code != "100000" {
		return "", fmt.Errorf("get chapter key error: %s", m.Tip)
	}
	if m.Data.Command == "" {
		return "", fmt.Errorf("get chapter key error: %s", "chapter key is empty")
	}
	return m.Data.Command, nil
}
func (api *API) GetChapterContentAPI(chapterId, chapterKey string) (*hbookermodel.ChapterInfo, error) {
	var content hbookermodel.Content
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId, "chapter_command": chapterKey}).Post(urlconstants.GetCptIfm)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&content)
	if content.Code != "100000" {
		return nil, fmt.Errorf("get chapter content error: %s", content.Tip)
	}
	if content.Data.ChapterInfo.TxtContent == "" {
		return nil, fmt.Errorf("get chapter content error: %s", "content is empty")
	}
	contentRaw, err := aesDecrypt(content.Data.ChapterInfo.TxtContent, chapterKey)
	if err != nil {
		return nil, err
	}
	content.Data.ChapterInfo.TxtContent = string(contentRaw)
	return &content.Data.ChapterInfo, nil
}

// Deprecated: MySignLogin is deprecated, hbooker has joined login verification, so this method is no longer available
func (api *API) MySignLogin(username, password, validate, challenge string) (*hbookermodel.LoginData, error) {
	var login hbookermodel.Login
	params := map[string]string{"login_name": username, "passwd": password}
	if validate != "" {
		params["geetest_seccode"] = validate + "|jordan"
		params["geetest_validate"] = validate
		params["geetest_challenge"] = challenge
	}
	res, err := api.DeleteValue("login_token").DeleteValue("account").
		HttpRequest.SetFormData(params).Post(urlconstants.MySignLogin)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&login)
	if login.Code != "100000" {
		return nil, fmt.Errorf("get login token error: %s", login.Tip)
	}
	if login.Data.LoginToken == "" {
		return nil, fmt.Errorf("get login token error: %s", "login token is empty")
	}

	return &login.Data, nil
}

func (api *API) GetBuyChapterAPI(chapterId string) (*hbookermodel.ContentBuy, error) {
	var m hbookermodel.ContentBuy
	res, err := api.HttpRequest.SetFormData(map[string]string{"chapter_id": chapterId, "shelf_id": ""}).Post(urlconstants.ChapterBuy)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&m)
	if m.Code != "100000" {
		return nil, fmt.Errorf("get buy chapter error: %s", m.Tip)
	}
	return &m, nil
}

func (api *API) GetAutoSignAPI(device string) (*hbookermodel.LoginData, error) {
	var m hbookermodel.Register
	if !checkDeviceRegex.MatchString(device) {
		return nil, fmt.Errorf("get auto sign error: %s", "device is not valid")
	}
	res, err := api.DeleteValue("login_token").DeleteValue("account").HttpRequest.
		SetFormData(map[string]string{"uuid": "android" + device, "gender": "1", "channel": "PCdownloadC"}).Post(urlconstants.SIGNUP)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&m)
	if m.Code != "100000" {
		return nil, fmt.Errorf("get auto sign error: %s", m.Tip)
	}
	return &m.Data, nil
}

func (api *API) GetUseGeetestAPI(loginName string) (*hbookermodel.GeetestData, error) {
	var geetest hbookermodel.Geetest
	res, err := api.HttpRequest.SetFormData(map[string]string{"login_name": loginName}).Post(urlconstants.UseGeetest)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&geetest)
	if geetest.Code != "100000" {
		return nil, fmt.Errorf("get geetest error: %s", geetest.Tip)
	}
	return &geetest.Data, nil
}

func (api *API) GetGeetestRegisterAPI(UserID string) (*hbookermodel.GeetestFirstRegisterStruct, error) {
	var geetestFirstRegister hbookermodel.GeetestFirstRegisterStruct
	res, err := api.HttpRequest.SetFormData(map[string]string{"user_id": UserID, "t": strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}).Post(urlconstants.GeetestRegister)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&geetestFirstRegister)
	if geetestFirstRegister.Challenge == "" {
		return nil, fmt.Errorf("get geetest register error: %s", "challenge is empty")
	}
	return &geetestFirstRegister, nil
}

func (api *API) GetBookShelfIndexesInfoAPI(shelfId string) (*hbookermodel.ShelfBook, error) {
	var bookList hbookermodel.ShelfBook
	res, err := api.HttpRequest.SetFormData(map[string]string{"shelf_id": shelfId, "direction": "prev", "last_mod_time": "0"}).Post(urlconstants.BookshelfGetShelfBookList)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&bookList)
	if bookList.Code != "100000" {
		return nil, fmt.Errorf("get book shelf indexes info error: %s", bookList.Tip)
	}
	if len(bookList.Data.BookList) == 0 {
		return nil, fmt.Errorf("get book shelf indexes info error: %s", "book list is empty")
	}
	return &bookList, nil
}

func (api *API) GetBookShelfInfoAPI() ([]hbookermodel.ShelfList, error) {
	var shelfList hbookermodel.Shelf
	res, err := api.HttpRequest.Post(urlconstants.BookshelfGetShelfList)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&shelfList)
	if shelfList.Code != "100000" {
		return nil, fmt.Errorf("get book shelf info error: %s", shelfList.Tip)
	}
	if len(shelfList.Data.ShelfList) == 0 {
		return nil, fmt.Errorf("get book shelf info error: %s", "shelf list is empty")
	}
	return shelfList.Data.ShelfList, nil
}

func (api *API) GetSearchBooksAPI(keyword string, page any) (*hbookermodel.Search, error) {
	var search hbookermodel.Search
	params := map[string]string{"count": "10", "page": fmt.Sprintf("%v", page), "category_index": "0", "key": keyword}
	res, err := api.HttpRequest.SetFormData(params).Post(urlconstants.BookcityGetFilterList)
	if err != nil {
		return nil, err
	}
	res.UnmarshalJson(&search)
	if search.Code != "100000" {
		return nil, fmt.Errorf("get search books error: %s", search.Tip)
	}
	if len(search.Data.BookList) == 0 {
		return nil, fmt.Errorf("get search books error: %s", "book list is empty")
	}
	return &search, nil
}
