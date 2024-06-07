package hbookermodel

type Search struct {
	Tip
	Data SearchData `json:"data"`
}

type SearchData struct {
	TagList []struct {
		TagName string `json:"tag_name"`
		Num     string `json:"num"`
	} `json:"tag_list"`
	BookList []BookInfo `json:"book_list"`
}

func (search *Search) EachBook(f func(int, BookInfo)) {
	if search.Data.BookList != nil {
		for i, book := range search.Data.BookList {
			f(i, book)
		}
	}

}
func (search *Search) GetBook(index int) *BookInfo {
	if search.Data.BookList != nil && index < len(search.Data.BookList) {
		return &search.Data.BookList[index]
	}
	return nil
}
