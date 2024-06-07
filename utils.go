package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
)

const (
	retryCount      = 5
	version         = "2.9.319"
	deviceToken     = "ciweimao_"
	androidApiKey   = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	userAgent       = "Android com.kuangxiangciweimao.novel "
	postContentType = "application/x-www-form-urlencoded"
	ivBase64        = "AAAAAAAAAAAAAAAAAAAAAA=="
)

var iv = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type continueFunction func(chapter hbookermodel.ChapterList) bool
type contentFunction func(chapter *hbookermodel.ChapterInfo)

type bookInfoFunction func(index int, bookInfo hbookermodel.BookInfo)
