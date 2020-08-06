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
	htmlData      map[string]interface{} //表示用のhtmlに渡すデータ型を定義
	MemoAppOutput struct {
		Results []*model.Memo
		Message string
	}

	memohandler struct {
		db repository.Database
	}
)

func ProvideMemohandler() *memohandler {
	return &memohandler{}
}

//引数はc(echo.Context型) 戻り値の型はerror
func (m *memohandler) MemoIndex(c echo.Context) error {
	memos, err := m.db.Get()
	if err != nil {
		c.Logger().Errorf("failed to select db request : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "メモが取得できませんでした"}) //構造体を渡すことによって、echoがJSONとして返す
	}
	// *********************追加*********************

	//index.htmlを返す。
	// return render(c, "src/views/index2.html", htmlData{"Memos": memos})
	return c.JSON(http.StatusOK,
		MemoAppOutput{
			Results: memos,
			Message: "取得OK"}) //構造体を渡すことによって、echoがJSONとして返す
	// *********************追加*********************

}

//引数はc(echo.Context型) 戻り値の型はerror
func (m *memohandler) MemoCreate(c echo.Context) error {
	var memo = &model.Memo{} //memoを定義

	err := c.Bind(memo) //フォームの内容を構造体に埋め込む
	if err != nil {
		c.Logger().Errorf("failed to bind : %v\n", err)
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "BadRequest"})
	}

	err = m.db.Set(c, memo) //repositoryを読み出して保存処理を実行
	if err != nil {
		c.Logger().Errorf("failed to create memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError, //サーバー内の処理でエラーが発生したら500エラーを返す
			MemoAppOutput{Message: "ServerError"})
	}

	return c.JSON(http.StatusOK,
		MemoAppOutput{Results: []*model.Memo{memo}, Message: "CreateSuccess"})
}

//削除機能
func (m *memohandler) MemoDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) //idを数値に変換してidに代入
	if err != nil {
		c.Logger().Errorf("failed to delete memo : %v\n", err)
		return c.JSON(http.StatusInternalServerError,
			MemoAppOutput{Message: "ServerErrror"})
	}
	//repositoryのメモ削除機のをを呼び出す
	if err := m.db.Delete(c, id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest,
			MemoAppOutput{Message: "BadRequest"})
	}
	// *********************追加*********************

	return c.JSON(http.StatusOK,
		MemoAppOutput{
			Message: fmt.Sprintf("Memo %d is deleted", id),
			Results: []*model.Memo{&model.Memo{ID: id}}})
	// *********************追加*********************

}
