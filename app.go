package hbookerLib

import (
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"github.com/AlexiaVeronica/input"
	"sync"
)

type APP struct {
	threadNum int
	bookInfo  *hbookermodel.BookInfo
	client    *Client
}

func (client *Client) APP() *APP {
	return &APP{client: client, threadNum: threadNum}
}
func (app *APP) SetThreadNum(threadNum int) *APP {
	app.threadNum = threadNum
	return app
}
func (app *APP) SetBookInfo(bookInfo *hbookermodel.BookInfo) *APP {
	app.bookInfo = bookInfo
	return app
}
func (app *APP) GetBookInfo() *hbookermodel.BookInfo {
	return app.bookInfo

}

func (app *APP) DownloadByChapterId(chapterId string) (*hbookermodel.ChapterInfo, error) {
	key, err := app.client.API().GetChapterCmd(chapterId)
	if err != nil {
		return nil, err
	}
	return app.client.API().GetCptIfm(chapterId, key.Data.Command)
}

func (app *APP) eachChapter(f func(hbookermodel.ChapterList)) *APP {
	divisionList, err := app.client.API().GetDivisionListByBookId(app.bookInfo.BookID)
	if err != nil {
		fmt.Println("get division list error:", err)
		return app
	}
	for _, division := range divisionList.Data.ChapterList {
		for _, chapter := range division.ChapterList {
			f(chapter)
		}
	}
	return app
}

func (app *APP) Download(continueFunc continueFunction, contentFunc contentFunction) *APP {
	var wg sync.WaitGroup
	if app.bookInfo == nil {
		fmt.Println("Please set book info first!")
		return app
	}
	ch := make(chan struct{}, app.threadNum)
	app.eachChapter(func(chapter hbookermodel.ChapterList) {
		wg.Add(1)
		ch <- struct{}{}
		go func(chapter hbookermodel.ChapterList) {
			defer func() {
				wg.Done()
				<-ch
			}()
			if continueFunc(app.bookInfo, chapter) {
				content, err := app.DownloadByChapterId(chapter.ChapterID)
				if err != nil {
					fmt.Println("get chapter content error:", err)
					return
				}
				contentFunc(app.bookInfo, content)
			}
		}(chapter)
	})
	wg.Wait()
	return app
}

func (app *APP) Search(keyword string, continueFunc continueFunction, contentFunc contentFunction) *APP {
	searchInfo, err := app.client.API().GetSearchBooksAPI(keyword, 0)
	if err != nil {
		fmt.Println("search failed!" + err.Error())
		return app
	}
	searchInfo.Each(func(index int, book hbookermodel.BookInfo) {
		fmt.Println("Index:", index, "\t\t\tBookName:", book.BookName)
	})
	app.bookInfo = searchInfo.GetBook(input.IntInput("Please input the index of the book you want to download"))
	return app.Download(continueFunc, contentFunc)
}

func (app *APP) Bookshelf(continueFunc continueFunction, contentFunc contentFunction) *APP {
	shelf, err := app.client.API().GetBookShelfInfoAPI()
	if err != nil {
		fmt.Println("get bookshelf error:", err)
		return app
	}
	for index, book := range shelf.Data.ShelfList {
		fmt.Println("Index:", index, "\t\t\tShelfName:", book.ShelfName, "\t\t\tShelfNum:", book.BookLimit)

	}
	bookshelf, err := app.client.API().GetBookcaseAPI(shelf.Data.ShelfList[input.IntInput("input the index of the bookshelf")].ShelfID)
	if err != nil {
		fmt.Println("get bookshelf error:", err)
		return app
	}
	bookshelf.Each(func(index int, book hbookermodel.BookInfo) {
		fmt.Println("Index:", index, "\t\t\tBookName:", book.BookName)
	})
	app.bookInfo = bookshelf.GetBook(input.IntInput("Please input the index of the book you want to download"))
	return app.Download(continueFunc, contentFunc)
}

func (app *APP) MergeText(f func(chapter hbookermodel.ChapterList)) {
	if app.bookInfo == nil {
		fmt.Println("Please set book info first!")
		return
	}
	app.eachChapter(f)
}
