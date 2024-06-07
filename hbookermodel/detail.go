package hbookermodel

type Detail struct {
	Tip
	Data struct {
		BookInfo    BookInfo   `json:"book_info" bson:"book_info"`
		RelatedList []BookInfo `json:"related_list" bson:"related_list"`
	} `json:"data" bson:"data"`
	ScrollChests []ScrollChest `json:"scroll_chests" bson:"scroll_chests"`
}

type BookInfo struct {
	BookID          string          `json:"book_id" bson:"book_id"`
	BookName        string          `json:"book_name" bson:"book_name"`
	Description     string          `json:"description" bson:"description"`
	BookSrc         string          `json:"book_src" bson:"book_src"`
	Tag             string          `json:"tag" bson:"tag"`
	TotalWordCount  string          `json:"total_word_count" bson:"total_word_count"`
	UpStatus        string          `json:"up_status" bson:"up_status"`
	UpdateStatus    string          `json:"update_status" bson:"update_status"`
	IsPaid          string          `json:"is_paid" bson:"is_paid"`
	Cover           string          `json:"cover" bson:"cover"`
	AuthorName      string          `json:"author_name" bson:"author_name"`
	Uptime          string          `json:"uptime" bson:"uptime"`
	Newtime         string          `json:"newtime" bson:"newtime"`
	ReviewAmount    string          `json:"review_amount" bson:"review_amount"`
	RewardAmount    string          `json:"reward_amount" bson:"reward_amount"`
	ChapterAmount   string          `json:"chapter_amount" bson:"chapter_amount"`
	LastChapterInfo LastChapterInfo `json:"last_chapter_info" bson:"last_chapter_info"`
	TagList         []TagList       `json:"tag_list" bson:"tag_list"`
	BookType        string          `json:"book_type" bson:"book_type"`
	TransverseCover string          `json:"transverse_cover" bson:"transverse_cover"`
}

type LastChapterInfo struct {
	ChapterID         string `json:"chapter_id" bson:"chapter_id"`
	BookID            string `json:"book_id" bson:"book_id"`
	ChapterIndex      string `json:"chapter_index" bson:"chapter_index"`
	ChapterTitle      string `json:"chapter_title" bson:"chapter_title"`
	Uptime            string `json:"uptime" bson:"uptime"`
	Mtime             string `json:"mtime" bson:"mtime"`
	RecommendBookInfo string `json:"recommend_book_info" bson:"recommend_book_info"`
}

type TagList struct {
	TagID   string `json:"tag_id" bson:"tag_id"`
	TagType string `json:"tag_type" bson:"tag_type"`
	TagName string `json:"tag_name" bson:"tag_name"`
}

type ScrollChest struct {
	ChestID     string `json:"chest_id" bson:"chest_id"`
	ReaderName  string `json:"reader_name" bson:"reader_name"`
	Gender      string `json:"gender" bson:"gender"`
	AvatarURL   string `json:"avatar_url" bson:"avatar_url"`
	BookName    string `json:"book_name" bson:"book_name"`
	Cost        int64  `json:"cost" bson:"cost"`
	ChestImgURL string `json:"chest_img_url" bson:"chest_img_url"`
	PropID      int64  `json:"prop_id" bson:"prop_id"`
	Content     string `json:"content" bson:"content"`
}
