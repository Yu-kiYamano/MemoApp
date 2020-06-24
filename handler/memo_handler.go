package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"memoapp/model"
	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

//type 識別子　型　で宣言
type (
	htmlData      map[string]interface{}
	MemoAppOutput struct {
		Memo    *model.Memo
		Message string
	}
)

//引数はc(echo.Context型) 戻り値の型はerror
func MemoIndex(c echo.Context) error {
	memos, err := repository.Getmemo()
	if err != nil {
		c.Logger().Errorf("failed to select db request : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"}) //JSONで返す理由:Messageを自由に設定する事ができるため
	}
	//index.htmlを返す。
	return render(c, "src/views/index.html", htmlData{"Memos": memos})
}

//引数はc(echo.Context型) 戻り値の型はerror
func MemoCreate(c echo.Context) error {
	var memo = &model.Memo{} //memoを定義

	err := c.Bind(memo) //フォームの内容を構造体に埋め込む
	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "error"})
	}

	res, err := repository.MemoCreate(c, memo) //repositoryを読み出して保存処理を実行
	if err != nil {
		c.Logger().Errorf("failed to create memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError, //サーバー内の処理でエラーが発生したら500エラーを返す
			MemoAppOutput{Message: "error"})
	}

	id, err := res.LastInsertId() //SQL実行結果から作成されたレコードのIDを取得する
	if err != nil {
		c.Logger().Errorf("failed to get ID : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}

	memo.SetId(int(id))
	return c.JSON(http.StatusOK,
		MemoAppOutput{Message: "success"})
}

//削除機能
func MemoDelete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id")) //idを数値に変換してidに代入
	if err != nil {
		c.Logger().Errorf("failed to delete memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "error"})
	}
	//repositoryのメモ削除機のをを呼び出す
	if err := repository.MemoDelete(c, id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "error"})
	}
	return c.JSON(http.StatusOK,
		fmt.Sprintf("Memo %d is deleted", id))
}
