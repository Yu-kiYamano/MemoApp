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
	db = connectDB()
	repository.SetDB(db)
	startServer()
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	//sqlx.Connectを使っても良い
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("データベースに接続しました")
	return db
}

func startServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.POST("/", handler.MemoCreate)
	e.GET("/", handler.MemoIndex)
	e.DELETE("/:id", handler.MemoDelete)
	e.Logger.Fatal(e.Start(":8080"))
	return e
}
