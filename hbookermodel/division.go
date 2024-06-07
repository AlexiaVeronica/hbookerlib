package hbookermodel

type NewVolumeList struct {
	Code string      `json:"code" bson:"code"`
	Tip  interface{} `json:"tip" bson:"tip"`
	Data struct {
		ChapterList []VolumeList `json:"chapter_list" bson:"chapter_list"`
	} `json:"data" bson:"data"`
	ScrollChests []interface{} `json:"scroll_chests" bson:"scroll_chests"`
}

type VolumeList struct {
	ChapterList     []ChapterList `json:"chapter_list" bson:"chapter_list"`
	MaxUpdateTime   string        `json:"max_update_time" bson:"max_update_time"`
	MaxChapterIndex string        `json:"max_chapter_index" bson:"max_chapter_index"`
	DivisionID      string        `json:"division_id" bson:"division_id"`
	DivisionIndex   string        `json:"division_index" bson:"division_index"`
	DivisionName    string        `json:"division_name" bson:"division_name"`
}

type ChapterList struct {
	ChapterID      string `json:"chapter_id" bson:"chapter_id"`
	ChapterIndex   string `json:"chapter_index" bson:"chapter_index"`
	ChapterTitle   string `json:"chapter_title" bson:"chapter_title"`
	WordCount      string `json:"word_count" bson:"word_count"`
	TsukkomiAmount string `json:"tsukkomi_amount" bson:"tsukkomi_amount"`
	IsPaid         string `json:"is_paid" bson:"is_paid"`
	Mtime          string `json:"mtime" bson:"mtime"`
	IsValid        string `json:"is_valid" bson:"is_valid"`
	AuthAccess     string `json:"auth_access" bson:"auth_access"`
}
