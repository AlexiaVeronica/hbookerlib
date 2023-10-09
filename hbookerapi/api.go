package hbookerapi

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type API struct {
	HttpClient HttpsClient
}

func (hbooker *API) GetBookInfo(bookId string) (*hbookermodel.BookInfo, error) {
	var book hbookermodel.Detail
	_, err := hbooker.HttpClient.Post(BOOK_GET_INFO_BY_ID, map[string]string{"book_id": bookId}, &book)
	if err != nil {
		return nil, err
	}
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

func (hbooker *API) GetDivisionListByBookId(bookId string) ([]hbookermodel.VolumeList, error) {
	var divisionList hbookermodel.NewVolumeList
	_, err := hbooker.HttpClient.Post(GET_DIVISION_LIST_NEW, map[string]string{"book_id": bookId}, &divisionList)
	if err != nil {
		return nil, err
	}
	if divisionList.Code != "100000" {
		return nil, fmt.Errorf("get division list error: %s", divisionList.Tip)
	}
	if len(divisionList.Data.ChapterList) == 0 {
		return nil, fmt.Errorf("get division list error: %s", "division list is empty")
	}
	return divisionList.Data.ChapterList, nil
}

func (hbooker *API) GetKetByChapterId(chapterId string) (string, error) {
	var m hbookermodel.ContentKey
	_, err := hbooker.HttpClient.Post(GET_CHAPTER_KEY, map[string]string{"chapter_id": chapterId}, &m)
	if err != nil {
		return "", err
	}
	if m.Code != "100000" {
		return "", fmt.Errorf("get chapter key error: %s", m.Tip)
	}
	if m.Data.Command == "" {
		return "", fmt.Errorf("get chapter key error: %s", "chapter key is empty")
	}
	return m.Data.Command, nil
}

func (hbooker *API) GetChapterContentAPI(chapterId, chapterKey string) (*hbookermodel.ChapterInfo, error) {
	var content hbookermodel.Content
	_, err := hbooker.HttpClient.Post(GET_CPT_IFM, map[string]string{"chapter_id": chapterId, "chapter_command": chapterKey}, &content)
	if err != nil {
		return nil, err
	}
	if content.Code != "100000" {
		return nil, fmt.Errorf("get chapter content error: %s", content.Tip)
	}
	if content.Data.ChapterInfo.TxtContent == "" {
		return nil, fmt.Errorf("get chapter content error: %s", "content is empty")
	}
	content.Data.ChapterInfo.TxtContent = string(HbookerDecode(content.Data.ChapterInfo.TxtContent, chapterKey))

	return &content.Data.ChapterInfo, nil
}

func (hbooker *API) GetLoginTokenAPI(username, password string) (*hbookermodel.Login, error) {
	var login hbookermodel.Login
	_, err := hbooker.HttpClient.Post(MY_SIGN_LOGIN, map[string]string{"login_name": username, "password": password}, &login)
	if err != nil {
		return nil, err
	}
	if login.Code != "100000" {
		return nil, fmt.Errorf("get login token error: %s", login.Tip)
	}
	return &login, nil
}
func (hbooker *API) GetBuyChapterAPI(chapterId, shelfId string) (*hbookermodel.ContentBuy, error) {
	var m hbookermodel.ContentBuy
	_, err := hbooker.HttpClient.Post(CHAPTER_BUY, map[string]string{"chapter_id": chapterId, "shelf_id": shelfId}, &m)
	if err != nil {
		return nil, err
	}
	if m.Code != "100000" {
		return nil, fmt.Errorf("get buy chapter error: %s", m.Tip)
	}
	return &m, nil
}

func (hbooker *API) GetAutoSignAPI(device string) (*hbookermodel.LoginData, error) {
	var m hbookermodel.Register
	checkDeviceRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	if !checkDeviceRegex.MatchString(device) {
		return nil, fmt.Errorf("get auto sign error: %s", "device is not valid")
	}
	params := map[string]string{"uuid": "android" + device, "gender": "1", "channel": "PCdownloadC"}
	_, err := hbooker.HttpClient.Post(SIGNUP, params, &m)
	if err != nil {
		return nil, err
	}
	if m.Code != "100000" {
		return nil, fmt.Errorf("get auto sign error: %s", m.Tip)
	}
	return &m.Data, nil
}

func (hbooker *API) GetUseGeetestAPI(loginName string) (*hbookermodel.Geetest, error) {
	var geetest hbookermodel.Geetest
	_, err := hbooker.HttpClient.Post(USE_GEETEST, map[string]string{"login_name": loginName}, &geetest)
	if err != nil {
		return nil, err
	}
	if geetest.Code != 100000 {
		return nil, fmt.Errorf("get geetest error: %s", geetest.Tip)
	}
	return &geetest, nil
}

func (hbooker *API) GetGeetestRegisterAPI(UserID string) (*hbookermodel.Challenge, error) {
	var challenge hbookermodel.Challenge
	_, err := hbooker.HttpClient.Post(GEETEST_REGISTER, map[string]string{"user_id": UserID, "t": strconv.FormatInt(time.Now().UnixNano()/1e6, 10)}, &challenge)
	if err != nil {
		return nil, err
	}
	if challenge.Challenge == "" {
		return nil, fmt.Errorf("get geetest register error: %s", "challenge is empty")
	}
	return &challenge, nil
}

func (hbooker *API) GetBookShelfIndexesInfoAPI(shelfId string) ([]hbookermodel.ShelfBookList, error) {
	var bookList hbookermodel.ShelfBook
	_, err := hbooker.HttpClient.Post(BOOKSHELF_GET_SHELF_BOOK_LIST, map[string]string{"shelf_id": shelfId, "direction": "prev", "last_mod_time": "0"}, &bookList)
	if err != nil {
		return nil, err
	}
	if bookList.Code != "100000" {
		return nil, fmt.Errorf("get book shelf indexes info error: %s", bookList.Tip)
	}
	if len(bookList.Data.BookList) == 0 {
		return nil, fmt.Errorf("get book shelf indexes info error: %s", "book list is empty")
	}
	return bookList.Data.BookList, nil
}

func (hbooker *API) GetBookShelfInfoAPI() ([]hbookermodel.ShelfList, error) {
	var shelfList hbookermodel.Shelf
	_, err := hbooker.HttpClient.Post(BOOKSHELF_GET_SHELF_LIST, map[string]string{}, &shelfList)
	if err != nil {
		return nil, err
	}
	if shelfList.Code != "100000" {
		return nil, fmt.Errorf("get book shelf info error: %s", shelfList.Tip)
	}
	if len(shelfList.Data.ShelfList) == 0 {
		return nil, fmt.Errorf("get book shelf info error: %s", "shelf list is empty")
	}
	return shelfList.Data.ShelfList, nil
}

func (hbooker *API) GetSearchBooksAPI(keyWord string, page int) ([]hbookermodel.BookInfo, error) {
	var search hbookermodel.Search
	_, err := hbooker.HttpClient.Post(BOOKCITY_GET_FILTER_LIST, map[string]string{"count": "10", "page": strconv.Itoa(page), "category_index": "0", "key": keyWord}, &search)
	if err != nil {
		return nil, err
	}
	if search.Code != "100000" {
		return nil, fmt.Errorf("get search books error: %s", search.Tip)
	}
	if len(search.Data.BookList) == 0 {
		return nil, fmt.Errorf("get search books error: %s", "book list is empty")
	}
	return search.Data.BookList, nil
}
