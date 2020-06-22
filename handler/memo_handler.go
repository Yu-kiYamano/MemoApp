package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"memoapp/model"
	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

type (
	htmlData      map[string]interface{}
	MemoAppOutput struct {
		Memo    *model.Memo
		Message string
	}
)

func MemoIndex(c echo.Context) error {
	memos, err := repository.Getmemo()
	if err != nil {
		c.Logger().Errorf("failed to select db request : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}

	return render(c, "src/views/index.html", htmlData{"Memos": memos})
}

func MemoCreate(c echo.Context) error {
	var memo = &model.Memo{}

	err := c.Bind(memo)
	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "error"})
	}

	res, err := repository.MemoCreate(c, memo)
	if err != nil {
		c.Logger().Errorf("failed to create memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.Logger().Errorf("failed to get ID : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}

	memo.SetId(int(id))
	// out.Memo = memo
	return c.JSON(http.StatusOK,
		MemoAppOutput{Message: "success"})
}

//削除機能
func MemoDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Errorf("failed to delete memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}
	if err := repository.MemoDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "error"})
	}
	return c.JSON(http.StatusOK,
		fmt.Sprintf("Memo %d is deleted", id))
}
