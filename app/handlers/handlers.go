package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{"Content": "Main page"})
}