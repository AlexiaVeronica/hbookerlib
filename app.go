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

func (app *APP) DownloadByChapterId(chapterId string) (*hbookermodel.ChapterInfo, error) {
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
				content, err := app.DownloadByChapterId(chapter.ChapterID)
				if err != nil {
					fmt.Println("get chapter content error:", err)
					return
				}
				f2(content)
			}
		}(chapter)
	})
	wg.Wait()
}

func (app *APP) Search(keyword string, f1 continueFunction, f2 contentFunction) {
	searchInfo, err := app.client.API().GetSearchBooksAPI(keyword, 0)
	if err != nil {
		fmt.Println("search failed!" + err.Error())
		return
	}
	searchInfo.EachBook(func(index int, book hbookermodel.BookInfo) {
		fmt.Println("Index:", index, "\t\t\tBookName:", book.BookName)
	})
	bookInfo := searchInfo.GetBook(GetUserInput("Please input the index of the book you want to download"))
	app.Download(bookInfo.BookID, f1, f2)
}

func (app *APP) Bookshelf(f1 continueFunction, f2 contentFunction) {
	shelf, err := app.client.API().GetBookShelfInfoAPI()
	if err != nil {
		fmt.Println("get bookshelf error:", err)
		return
	}
	for index, book := range shelf {
		fmt.Println("Index:", index, "\t\t\tShelfName:", book.ShelfName, "\t\t\tShelfNum:", book.BookLimit)

	}
	bookshelf, err := app.client.API().GetBookShelfIndexesInfoAPI(shelf[GetUserInput("Please input the index of the bookshelf you want to download")].ShelfID)
	if err != nil {
		fmt.Println("get bookshelf error:", err)
		return
	}
	bookshelf.EachBook(func(index int, book hbookermodel.BookInfo) {
		fmt.Println("Index:", index, "\t\t\tBookName:", book.BookName)
	})
	bookInfo := bookshelf.GetBook(GetUserInput("Please input the index of the book you want to download"))
	app.Download(bookInfo.BookID, f1, f2)

}
