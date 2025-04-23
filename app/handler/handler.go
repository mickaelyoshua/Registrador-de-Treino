package handler

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/view"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/helper"
)

func Render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}
func HandleRenderError(err error) {
	if err != nil {
		log.Fatalf("Could not render template: %v", err)
	}
}

// Main page
func Index(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Index())
	HandleRenderError(err)
}

// Authentication
func Register(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Register())
	HandleRenderError(err)
}

func ConfirmPass(ctx *gin.Context) {
	pass := ctx.Request.FormValue("password")
	confirm := ctx.Request.FormValue("confirmPassword")

	if !helper.ValidatePassword(pass, confirm) {
		ctx.String(http.StatusBadRequest, "Senhas estão diferentes")
	} else {
		ctx.Status(http.StatusOK)
	}
}