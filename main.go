package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// const tmplPath = "src/view"

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	e.GET("/", Index)
	e.GET("/:id", Show)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
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

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/styles", "src/styles")

	return e
}

func Index(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "index",
	}
	return render(c, "src/views/index.html", data)
}

func Show(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Show",
	}
	return render(c, "src/views/show.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
