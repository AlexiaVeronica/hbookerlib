package hbookermodel

type ReaderInfo struct {
	ReaderId       string        `json:"reader_id"`
	Account        string        `json:"account"`
	IsBind         string        `json:"is_bind"`
	IsBindQq       string        `json:"is_bind_qq"`
	IsBindWeixin   string        `json:"is_bind_weixin"`
	IsBindHuawei   string        `json:"is_bind_huawei"`
	IsBindApple    string        `json:"is_bind_apple"`
	PhoneNum       string        `json:"phone_num"`
	PhoneCrypto    string        `json:"phone_crypto"`
	MobileVal      string        `json:"mobileVal"`
	Email          string        `json:"email"`
	License        string        `json:"license"`
	ReaderName     string        `json:"reader_name"`
	AvatarUrl      string        `json:"avatar_url"`
	AvatarThumbUrl string        `json:"avatar_thumb_url"`
	BaseStatus     string        `json:"base_status"`
	ExpLv          string        `json:"exp_lv"`
	ExpValue       string        `json:"exp_value"`
	Gender         string        `json:"gender"`
	VipLv          string        `json:"vip_lv"`
	VipValue       string        `json:"vip_value"`
	IsAuthor       string        `json:"is_author"`
	IsUploader     string        `json:"is_uploader"`
	BookAge        string        `json:"book_age"`
	CategoryPrefer []interface{} `json:"category_prefer"`
	UsedDecoration []struct {
		DecorationType     string `json:"decoration_type"`
		DecorationUrl      string `json:"decoration_url"`
		DecorationId       string `json:"decoration_id"`
		ReaderDecorationId string `json:"reader_decoration_id"`
	} `json:"used_decoration"`
	Rank         string `json:"rank"`
	FirstLoginIp string `json:"first_login_ip"`
	Ctime        string `json:"ctime"`
}

type PropInfo struct {
	RESTGiftHlb     string `json:"rest_gift_hlb"`
	RESTHlb         string `json:"rest_hlb"`
	RESTYp          string `json:"rest_yp"`
	RESTRecommend   string `json:"rest_recommend"`
	RESTTotalBlade  string `json:"rest_total_blade"`
	RESTMonthBlade  string `json:"rest_month_blade"`
	RESTTotal100    string `json:"rest_total_100"`
	RESTTotal588    string `json:"rest_total_588"`
	RESTTotal1688   string `json:"rest_total_1688"`
	RESTTotal5000   string `json:"rest_total_5000"`
	RESTTotal10000  string `json:"rest_total_10000"`
	RESTTotal100000 string `json:"rest_total_100000"`
	RESTTotal50000  string `json:"rest_total_50000"`
	RESTTotal160000 string `json:"rest_total_160000"`
}
