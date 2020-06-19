package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"memoapp/model"
	"memoapp/repository"

	"github.com/labstack/echo/v4"
)

func MemoIndex(c echo.Context) error {
	// リポジトリの処理を呼び出して記事の一覧データを取得します。
	memos, err := repository.MemoListByCursor(0)

	// エラーが発生した場合
	if err != nil {
		// エラー内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// クライアントにステータスコード 500 でレスポンスを返します。
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータを map に格納します。
	data := map[string]interface{}{
		"Memos": memos,
	}

	// テンプレートファイルとデータを指定して HTML を生成し、クライアントに返却します
	return render(c, "src/views/index.html", data)
}

// ArticleCreateOutput ...
type MemoCreateOutput struct {
	Memo             *model.Memo
	Message          string
	ValidationErrors []string
}

// ArticleCreate ...
func MemoCreate(c echo.Context) error {
	// 送信されてくるフォームの内容を格納する構造体を宣言します。
	var memo model.Memo

	// レスポンスとして返却する構造体を宣言します。
	var out MemoCreateOutput

	// フォームの内容を構造体に埋め込みます。
	if err := c.Bind(&memo); err != nil {
		// エラーの内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// リクエストの解釈に失敗した場合は 400 エラーを返却します。
		return c.JSON(http.StatusBadRequest, out)

	}

	// repository を呼び出して保存処理を実行します。
	res, err := repository.MemoCreate(&memo)
	if err != nil {
		// エラーの内容をサーバーのログに出力します。
		c.Logger().Error(err.Error())

		// サーバー内の処理でエラーが発生した場合は 500 エラーを返却します。
		return c.JSON(http.StatusInternalServerError, out)
	}

	// SQL 実行結果から作成されたレコードの ID を取得します。
	id, _ := res.LastInsertId()

	// 構造体に ID をセットします。
	memo.ID = int(id)

	// レスポンスの構造体に保存した記事のデータを格納します。
	out.Memo = &memo

	// 処理成功時はステータスコード 200 でレスポンスを返却します。
	return c.JSON(http.StatusOK, out)
}

func MemoDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := repository.MemoDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Memo %d is deleted", id))
}
