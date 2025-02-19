package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Count struct {
	Count int
}

func main() {
	count := Count{Count: 0}
	router := gin.Default()

	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", count) // Using index and not index.html because is identifying the block, not the file
	})

	router.POST("/count", func(c *gin.Context) {
		count.Count++
		c.HTML(http.StatusOK, "count", count) // Using index and not index.html because is identifying the block, not the file
	})

	router.Run(":8080")
}