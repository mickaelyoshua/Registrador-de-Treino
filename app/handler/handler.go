package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/Registrador-de-Treino/db"
	"github.com/mickaelyoshua/Registrador-de-Treino/util"
	"github.com/mickaelyoshua/Registrador-de-Treino/model"
	"github.com/mickaelyoshua/Registrador-de-Treino/view"
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
func Hi(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		log.Fatalf("Error getting cookie: \n%v", err)
	}
	err = Render(ctx, http.StatusOK, view.Hi(username))
	HandleRenderError(err)
}

// Authentication
func RegisterView(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Register())
	HandleRenderError(err)
}
func LoginView(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Login())
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
	password, err := util.HashPassword(ctx.Request.FormValue("password"))
	if err != nil {
		log.Fatalf("Error hashing password: \n%v", err)
	}

	created := time.Now()
	user := model.NewUser(username, email, password, created, created)
	err = user.Save(client)
	if err != nil {
		log.Fatalf("Error saving User: \n%v", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func Login(ctx *gin.Context) {
	client, err := db.GetClient()
	if err != nil {
		log.Fatalf("Error getting client from MongoDB: \n%v", err)
		return
	}

	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	user, err := model.FindUser(client, map[string]string{"email": email})
	if err != nil {
		log.Fatalf("Error finding user by email: \n%v", err)
		ctx.String(http.StatusBadRequest, "Usuário não encontrado")
		return
	}

	if user == nil {
		ctx.String(http.StatusBadRequest, "Usuário não encontrado")
		return
	}

	if !util.CheckPasswordHash(password, user.Password) {
		ctx.String(http.StatusBadRequest, "Senha incorreta")
		return
	}

	ctx.SetCookie("username", user.Username, 3600, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/hi")
}

func ConfirmPass(ctx *gin.Context) {
	pass := ctx.Request.FormValue("password")
	confirm := ctx.Request.FormValue("confirmPassword")
	fmt.Println(pass, confirm)

	if !util.ValidatePassword(pass, confirm) {
		ctx.String(http.StatusBadRequest, "Senhas estão diferentes")
	} else {
		ctx.Status(http.StatusOK)
	}
}