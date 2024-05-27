package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"sync"
)

type APP struct {
	threadNum int
	client    *Client
}

func (client *Client) APP() *APP {
	return &APP{client: client, threadNum: 32}
}
func (app *APP) SetThreadNum(threadNum int) *APP {
	app.threadNum = threadNum
	return app
}

func (app *APP) NewGetContent(chapterId string) (*hbookermodel.ChapterInfo, error) {
	key, err := app.client.API().GetChapterKey(chapterId)
	if err != nil {
		return nil, err
	}
	return app.client.API().GetChapterContentAPI(chapterId, key)
}

func (app *APP) EachChapter(bookId string, f func(hbookermodel.ChapterList)) {
	divisionList, err := app.client.API().GetDivisionListByBookId(bookId)
	if err != nil {
		fmt.Println("get division list error:", err)
		return
	}
	for _, division := range divisionList {
		for _, chapter := range division.ChapterList {
			f(chapter)
		}
	}
}

func (app *APP) Download(bookId string, f1 continueFunction, f2 contentFunction) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, app.threadNum)
	app.EachChapter(bookId, func(chapter hbookermodel.ChapterList) {
		wg.Add(1)
		ch <- struct{}{}
		go func(chapter hbookermodel.ChapterList) {
			defer func() {
				wg.Done()
				<-ch
			}()
			if f1(chapter) {
				content, err := app.NewGetContent(chapter.ChapterID)
				if err != nil {
					fmt.Println("get chapter content error:", err)
					return
				}
				f2(chapter, content.TxtContent)
			}
		}(chapter)
	})
	wg.Wait()
}
