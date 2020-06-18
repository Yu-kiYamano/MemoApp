package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	// 記事データの一覧を取得する
	memos, err := repository.MemoList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message": "Indexページ",
		"Now":     time.Now(),
		"memos":   memos, // 記事データをテンプレートエンジンに渡す
	}
	return render(c, "src/views/index.html", data)
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"Message": "Showページ",
		"Now":     time.Now(),
		"id":      id,
	}
	return render(c, "src/views/show.html", data)
}
