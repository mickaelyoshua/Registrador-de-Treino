package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autheticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}