package hbookerLib

import (
	"bufio"
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	retryCount      = 5
	version         = "2.9.290"
	deviceToken     = "ciweimao_"
	androidApiKey   = "zG2nSeEfSHfvTCHy5LCcqtBbQehKNLXn"
	userAgent       = "Android com.kuangxiangciweimao.novel "
	postContentType = "application/x-www-form-urlencoded"
)

var iv = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var checkDeviceRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

type continueFunction func(chapter hbookermodel.ChapterList) bool
type contentFunction func(chapter *hbookermodel.ChapterInfo)

type bookInfoFunction func(index int, bookInfo hbookermodel.BookInfo)

func GetUserInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		return number
	}
}
