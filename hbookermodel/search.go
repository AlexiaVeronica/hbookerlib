package hbookermodel

type Search struct {
	Code string      `json:"code"`
	Data SearchData  `json:"data"`
	Tip  interface{} `json:"tip"`
}

type SearchData struct {
	TagList []struct {
		TagName string `json:"tag_name"`
		Num     string `json:"num"`
	} `json:"tag_list"`
	BookList []BookInfo `json:"book_list"`
}
