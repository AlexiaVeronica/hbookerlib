package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"regexp"
)

const (
	version         = "2.9.290"
	deviceToken     = "ciweimao_"
	androidApiKey   = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	retryCount      = 5
	userAgent       = "Android com.kuangxiangciweimao.novel "
	postContentType = "application/x-www-form-urlencoded"
)

var iv = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var checkDeviceRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

type continueFunction func(chapter hbookermodel.ChapterList) bool
type contentFunction func(chapter hbookermodel.ChapterList, content string)
