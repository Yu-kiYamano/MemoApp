package handler

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) { //テンプレートが正常に解析できなかったらパニック
	return pongo2.Must(pongo2.FromCache(file)).ExecuteBytes(data) //htmlを生成。作られたhtmlはバイトデータとして呼び出し元に返される
}

//引数はc(echo.COntext型),file(string型),data(stringのmap) 返り値はerror
func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data) //htmlBlobを呼び出してfile(生成されたhtml)とバイトデータを受けとる

	if err != nil {
		return c.NoContent(http.StatusInternalServerError) //エラーが起きたらからのbodyとエラーコードを返す
	}
	return c.HTMLBlob(http.StatusOK, b)
}
