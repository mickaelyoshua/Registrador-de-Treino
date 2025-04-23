package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/db"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/helper"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/model"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/view"
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
func RegisterView(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Register())
	HandleRenderError(err)
}

func Register(ctx *gin.Context) {
	client, err := db.GetClient()
	if err != nil {
		log.Fatalf("Error getting client from MongoDB: \n%v", err)
		return
	}

	username := ctx.Request.FormValue("username")
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	user := model.NewUser(username, email, password, time.Now(), time.Now())
	err = user.Save(client)
	if err != nil {
		log.Fatalf("Error saving User: \n%v", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func ConfirmPass(ctx *gin.Context) {
	pass := ctx.Request.FormValue("password")
	confirm := ctx.Request.FormValue("confirmPassword")
	fmt.Println(pass, confirm)

	if !helper.ValidatePassword(pass, confirm) {
		ctx.String(http.StatusBadRequest, "Senhas estão diferentes")
	} else {
		ctx.Status(http.StatusOK)
	}
}