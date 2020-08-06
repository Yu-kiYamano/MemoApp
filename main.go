package main

import (
	"memoapp/handler"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

func main() {

	startServer() //サーバースタート
}

func startServer() *echo.Echo {

	e := echo.New() //echoインスタンス作成
	// dsn := os.Getenv("DSN")            //.envrcのDSNを取得してdsnに代入(dsnとはプログラム側が捜査対象のdbを指定するための識別子)
	// db, err := sqlx.Open("mysql", dsn) //("mysql"(ドライバ名),dsn(dsnの名前(26行目で定義))　*sql.DB(つまりdb)を返す )
	// if err != nil {
	// 	e.Logger.Errorf("データベース接続に失敗しました。: %v\n", err)
	// }
	// if err := db.Ping(); err != nil { //Pingとは対処のコンピュータとネットワークで繋がっているかを確認する時に使うもの
	// 	e.Logger.Errorf("確認できません: %v\n", err)
	// }
	// log.Println("データベースに接続しました")

	// defer db.Close() //startServer関数が終了する際に実行されるようにする為、deferを記述

	// echoログのフォーマット
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	})

	hdlr := handler.ProvideMemohandler()
	//middlewareを登録
	e.Use(
		logger,
		middleware.Recover(), //パニックから回復させるためのmiddleware
		middleware.Logger(),  //各HTTP リクエストに関する情報をログに記録するためのmiddleware
		middleware.Gzip(),    //gzip圧縮スキームを使用してHTTPレスポンスを圧縮するためのmiddleware
		hdlr.CheckCache(),
	)

	e.POST("/", hdlr.MemoCreate)
	e.GET("/list", hdlr.MemoIndex)

	// インデックス画面を表示
	e.GET("/", index)

	e.DELETE("/:id", hdlr.MemoDelete)
	e.Logger.Fatal(e.Start(":8080"))
	return e

}

// *********************追加*********************
func index(c echo.Context) error {
	return handler.Render(c, "src/views/index.html", nil)
}

func logFormat() string {
	var format string
	format += "\n[  echo ]"
	format += "time:${time_rfc3339}\n"
	format += "- method:${method}\t"
	format += "status:${status}\n"
	format += "- error:${error}\n"
	format += "- path:${path}\t"
	format += "uri:${uri}\t"
	format += "host:${host}\t"
	format += "remote_ip:${remote_ip}\n"
	format += "- bytes_in:${bytes_in}\t"
	format += "bytes_out:${bytes_out}\n"
	format += "- latency:${latency}\t"
	format += "latency_human:${latency_human}\n\n"
	// format += "forwardedfor:${header:x-forwarded-for}\n"
	// format += "referer:${referer}\n"
	// format += "user_agent:${user_agent}\n"
	// format += "request_id:${id}\n"

	return format
}

// *********************追加*********************
