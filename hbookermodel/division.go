package hbookermodel

type NewVolumeList struct {
	Code string      `json:"code"`
	Tip  interface{} `json:"tip"`
	Data struct {
		ChapterList []VolumeList `json:"chapter_list"`
	} `json:"data"`
	ScrollChests []interface{} `json:"scroll_chests"`
}

type VolumeList struct {
	ChapterList     []ChapterList `json:"chapter_list"`
	MaxUpdateTime   string        `json:"max_update_time"`
	MaxChapterIndex string        `json:"max_chapter_index"`
	DivisionID      string        `json:"division_id"`
	DivisionIndex   string        `json:"division_index"`
	DivisionName    string        `json:"division_name"`
}
type ChapterList struct {
	ChapterID      string `json:"chapter_id"`
	ChapterIndex   string `json:"chapter_index"`
	ChapterTitle   string `json:"chapter_title"`
	WordCount      string `json:"word_count"`
	TsukkomiAmount string `json:"tsukkomi_amount"`
	IsPaid         string `json:"is_paid"`
	Mtime          string `json:"mtime"`
	IsValid        string `json:"is_valid"`
	AuthAccess     string `json:"auth_access"`
}
