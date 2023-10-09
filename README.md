# HbookerLib

The `HbookerLib` package provides a client library for interacting with the Hbooker API. It allows you to perform various operations such as retrieving book information, getting chapter content, buying chapters, managing bookshelves, and more.

## Installation

To install the `HbookerLib` package, use the following command:

```shell
go get github.com/AlexiaVeronica/hbookerLib
```

## Usage

Import the `HbookerLib` package in your Go code:

```go
import "github.com/AlexiaVeronica/hbookerLib"
```

Create a new client by calling the `NewClient` function:

```go
client := HbookerLib.NewClient()
```

You can customize the client by using various options:

```go
client := HbookerLib.NewClient(
    HbookerLib.WithLoginToken("your_login_token"),
    HbookerLib.WithAccount("your_account"),
    HbookerLib.WithVersion("your_version"),
    HbookerLib.WithDebug(),
    HbookerLib.WithOutputDebug(),
    HbookerLib.WithProxyURLArray([]string{"proxy_url_1", "proxy_url_2"}),
    HbookerLib.WithProxyURL("proxy_url"),
    HbookerLib.WithAPIBaseURL("your_api_base_url"),
    HbookerLib.WithUserAgent("your_user_agent"),
    HbookerLib.WithAndroidApiKey("your_android_api_key"),
)
```

### Available Options

- `WithLoginToken(loginToken string) Options`: Sets the login token for the client.
- `WithAccount(account string) Options`: Sets the account for the client.
- `WithVersion(version string) Options`: Sets the version for the client.
- `WithDebug() Options`: Enables debug mode for the client.
- `WithOutputDebug() Options`: Enables output debug mode for the client.
- `WithProxyURLArray(proxyURLArray []string) Options`: Sets the proxy URL array for the client.
- `WithProxyURL(proxyURL string) Options`: Sets the proxy URL for the client.
- `WithAPIBaseURL(apiBaseURL string) Options`: Sets the API base URL for the client.
- `WithUserAgent(userAgent string) Options`: Sets the user agent for the client.
- `WithAndroidApiKey(androidApiKey string) Options`: Sets the Android API key for the client.
 

## API Reference

The `HbookerLib` package provides the following methods:

- `GetBookInfo(bookId string) (*models.BookInfo, error)`: Retrieves the information of a book by its ID.
- `GetDivisionListByBookId(bookId string) ([]models.VolumeList, error)`: Retrieves the division list of a book by its ID.
- `GetKetByChapterId(chapterId string) (string, error)`: Retrieves the key of a chapter by its ID.
- `GetChapterContentAPI(chapterId, chapterKey string) (*models.ChapterInfo, error)`: Retrieves the content of a chapter by its ID and key.
- `GetLoginTokenAPI(username, password string) (*models.Login, error)`: Retrieves the login token for the specified username and password.
- `GetBuyChapterAPI(chapterId, shelfId string) (*models.ContentBuy, error)`: Buys a chapter with the specified ID and adds it to the specified shelf.
- `GetAutoSignAPI(device string) (*models.LoginData, error)`: Retrieves the auto sign information for the specified device.
- `GetUseGeetestAPI(loginName string) (*models.Geetest, error)`: Retrieves the geetest information for the specified login name.
- `GetGeetestRegisterAPI(UserID string) (*models.Challenge, error)`: Retrieves the geetest register information for the specified user ID.
- `GetBookShelfIndexesInfoAPI(shelfId string) ([]models.ShelfBookList, error)`: Retrieves the bookshelf indexes information for the specified shelf ID.
- `GetBookShelfInfoAPI() ([]models.ShelfList, error)`: Retrieves the bookshelf information.
- `GetSearchBooksAPI(keyWord string, page int) ([]models.BookInfo, error)`: Retrieves the search results for the specified keyword and page number.

## License

This package is licensed under the [MIT License](https://github.com/AlexiaVeronica/hbookerLib/licenses/MIT).
