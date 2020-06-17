package main

<<<<<<< HEAD
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", Index)
	router.Run(":8080")
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
=======
import(
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template"

var e = createMux()

func main(){
	e.GET("/",articleIndex)
	e.Logger.Fatal(e.Start("8080"))
}

func createMux() *echo.Echo{
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error{
	data := map[string]interface{}{
		"Message": "Hello,World",
		"Now": time.Now(),
	}
	return render(c echo.Context,file string,data map[string]interface{}) error{
		b, err := htmlBlob(file,data)
		if err != nil{
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.HTMLBlob(http.statusOK,b)
	}
}
>>>>>>> origin/backend
