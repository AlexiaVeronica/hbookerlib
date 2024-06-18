package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
)

const (
	retryCount      = 5
	version         = "2.9.319"
	versionIos      = "3.2.710"
	deviceToken     = "ciweimao_"
	deviceIosToken  = "iPhone-"
	threadNum       = 32
	androidApiKey   = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	userAgent       = "Android com.kuangxiangciweimao.novel "
	postContentType = "application/x-www-form-urlencoded"
	ivBase64        = "AAAAAAAAAAAAAAAAAAAAAA=="
)

type continueFunction func(bookInfo *hbookermodel.BookInfo, chapter hbookermodel.ChapterList) bool
type contentFunction func(bookInfo *hbookermodel.BookInfo, chapter *hbookermodel.ChapterInfo)
