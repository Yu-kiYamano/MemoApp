package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"memoapp/model"
	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	//一覧を取得
	memos, err := repository.MemoList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message": "Indexページ",
		"Now":     time.Now(),
		"memos":   memos, // テンプレートエンジンに渡す
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

type MemoCreateOutput struct {
	Memo             *model.Memo
	Message          string
	ValidationErrors []string
}

func MemoCreate(c echo.Context) error {
	var memo model.Memo

	var out MemoCreateOutput

	if err := c.Bind(&memo); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, out)
	}

	res, err := repository.MemoCreate(&memo)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, out)
	}

	id, _ := res.LastInsertId()

	memo.ID = int(id)
	out.Memo = &memo

	return c.JSON(http.StatusOK, out)
}
