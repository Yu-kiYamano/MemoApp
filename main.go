package main

import (
	"log"
	"os"

	"memoapp/handler"
	"memoapp/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

func main() {

	startServer()
}

func startServer() *echo.Echo {

	e := echo.New()
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	repository.SetDB(db)
	if err != nil {
		e.Logger.Errorf("データベース接続に失敗しました。: %v\n", err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Errorf("確認できません: %v\n", err)
	}
	log.Println("データベースに接続しました")

	//middlewareを登録
	e.Use(
		middleware.Recover(), //パニックから回復させるためのmiddleware
		middleware.Logger(),  //各HTTP リクエストに関する情報をログに記録するためのmiddleware
		middleware.Gzip(),    //gzip圧縮スキームを使用してHTTPレスポンスを圧縮するためのmiddleware
	)

	e.POST("/", handler.MemoCreate)
	e.GET("/", handler.MemoIndex)
	e.DELETE("/:id", handler.MemoDelete)
	e.Logger.Fatal(e.Start(":8080"))
	return e

}
