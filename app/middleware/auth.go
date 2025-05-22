package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/Registrador-de-Treino/util"
)

func Authenticate(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	// Validate the token
	retrievedToken, err := util.ValidateToken(token)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	// Check if the token is expired
	if retrievedToken["expiration"].(int64) < time.Now().Unix() {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
}