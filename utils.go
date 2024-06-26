package hbookerLib

import (
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
)

const (
	retryCount       = 5
	versionAndroid   = "2.9.328"
	versionIos       = "3.2.710"
	deviceToken      = "ciweimao_"
	deviceIosToken   = "iPhone-"
	threadNum        = 32
	apiKey           = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	androidUserAgent = "Android com.kuangxiangciweimao.novel %v"
	iosUserAgent     = "HappyBook/%v (iPhone; iOS 14.5.1; Scale/2.00)"
	postContentType  = "application/x-www-form-urlencoded"
	ivBase64         = "AAAAAAAAAAAAAAAAAAAAAA=="

	hmacKey              = "a90f3731745f1c30ee77cb13fc00005a"
	androidSignaturesKey = "CkMxWNB666"
	IosSignaturesKey     = "kuangxiang.HappyBook"
	publicIOSKey         = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCohMLlejVLZvmFh/XFG2N5YKAjCeU08hiWUXGTUztFcUnYYhv2J1FknW/FuinK+ojveEYTNpHeXvXBjc7PXVGYLzCt+B4XW7zheehTcE8Wut3IzJd8rnIUbNpqLgqe6Ttu/X46E8wI8Xnkxlluh0wPRPIu+MmqyS1k6+2A6m/tQIDAQAB`
)

type continueFunction func(bookInfo *hbookermodel.BookInfo, chapter hbookermodel.ChapterList) bool
type contentFunction func(bookInfo *hbookermodel.BookInfo, chapter *hbookermodel.ChapterInfo)
