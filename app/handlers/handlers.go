package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/models"
)

func DefaultHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"Content": "Main page"})
}

func Register(ctx *gin.Context) {
	email := ctx.Request.FormValue("email")
	pass := ctx.Request.FormValue("password")
	u := models.User{Email: email, Password: pass}

	id, err := u.Save()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error saving user\nError: %s", err.Error())
	} else {
		ctx.String(http.StatusOK, "Saved with id %v", id)
	}
}

func Login(ctx *gin.Context) {

}