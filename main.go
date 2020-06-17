package main

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
