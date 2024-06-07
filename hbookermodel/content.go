package hbookermodel

type Content struct {
	Code string      `json:"code" bson:"code"`
	Tip  interface{} `json:"tip" bson:"tip"`
	Data struct {
		ChapterInfo ChapterInfo `json:"chapter_info" bson:"chapter_info"`
	} `json:"data" bson:"data"`
}

type ContentKey struct {
	Code string      `json:"code" bson:"code"`
	Tip  interface{} `json:"tip" bson:"tip"`
	Data struct {
		Command string `json:"command" bson:"command"`
	} `json:"data" bson:"data"`
}

type ChapterInfo struct {
	ChapterID         string `json:"chapter_id" bson:"chapter_id"`
	BookID            string `json:"book_id" bson:"book_id"`
	DivisionID        string `json:"division_id" bson:"division_id"`
	UnitHlb           string `json:"unit_hlb" bson:"unit_hlb"`
	ChapterIndex      string `json:"chapter_index" bson:"chapter_index"`
	ChapterTitle      string `json:"chapter_title" bson:"chapter_title"`
	AuthorSay         string `json:"author_say" bson:"author_say"`
	WordCount         string `json:"word_count" bson:"word_count"`
	Discount          string `json:"discount" bson:"discount"`
	IsPaid            string `json:"is_paid" bson:"is_paid"`
	AuthAccess        string `json:"auth_access" bson:"auth_access"`
	BuyAmount         string `json:"buy_amount" bson:"buy_amount"`
	TsukkomiAmount    string `json:"tsukkomi_amount" bson:"tsukkomi_amount"`
	TotalHlb          string `json:"total_hlb" bson:"total_hlb"`
	Uptime            string `json:"uptime" bson:"uptime"`
	Mtime             string `json:"mtime" bson:"mtime"`
	RecommendBookInfo string `json:"recommend_book_info" bson:"recommend_book_info"`
	TxtContent        string `json:"txt_content" bson:"txt_content"`
}

type ContentBuyData struct {
	ReaderInfo ReaderInfo `json:"reader_info" bson:"reader_info"`
	PropInfo   PropInfo   `json:"prop_info" bson:"prop_info"`
	BookInfo   struct {
		IsBuy       string      `json:"is_buy" bson:"is_buy"`
		BookInfo    BookInfo    `json:"book_info" bson:"book_info"`
		ShelfId     string      `json:"shelf_id" bson:"shelf_id"`
		ChapterInfo ChapterInfo `json:"chapter_info" bson:"chapter_info"`
	} `json:"book_info" bson:"book_info"`
}

type ContentBuy struct {
	Code         string         `json:"code" bson:"code"`
	Tip          string         `json:"tip" bson:"tip"`
	Data         ContentBuyData `json:"data" bson:"data"`
	ScrollChests ScrollChest    `json:"scroll_chests" bson:"scroll_chests"`
}
