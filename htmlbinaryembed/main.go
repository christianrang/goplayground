package main

import (
	_ "embed"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
var indexHtml string

func main() {
	fmt.Println(indexHtml)
	router := gin.Default()
	router.GET("/", CachedPageHandler)
	router.Run("localhost:8080")
}

func CachedPageHandler(c *gin.Context) {
	//Write your 200 header status (or other status codes, but only WriteHeader once)
	c.Writer.WriteHeader(http.StatusOK)
	//Convert your cached html string to byte array
	if _, err := c.Writer.Write([]byte(indexHtml)); err != nil {
		log.Fatalln("error writing html: ", err)
	}
	return
}
