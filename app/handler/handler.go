package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	token, err := ctx.Cookie("token")
	if err != nil {
		log.Println("Token not found or invalid:", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	retrievedToken, err := util.ValidateToken(token)
	if err != nil {
		log.Println("Invalid or expired token:", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	client, err := db.GetClient()
	if err != nil {
		log.Fatalf("Error getting client from MongoDB: \n%v", err)
		return
	}

	// Convert the ID from the token to ObjectID
	objectID, err := primitive.ObjectIDFromHex(retrievedToken["id"].(string))
	if err != nil {
		log.Fatalf("Error converting ID to ObjectID: \n%v", err)
		return
	}

	user, err := model.FindUser(client, bson.M{"_id": objectID})
	if err != nil {
		log.Fatalf("Error finding user by ID: \n%v", err)
		return
	}

	err = Render(ctx, http.StatusOK, view.Index(user))
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

	token, err := util.GenerateToken(email, user.Id) // Pass ObjectID directly
	if err != nil {
		log.Fatalf("Error generating token: \n%v", err)
		ctx.String(http.StatusBadRequest, "Erro ao gerar token")
		return
	}
	ctx.SetCookie("token", token, int(2*time.Hour.Seconds()), "/", "", false, true)

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

	if user.Username == "" {
		ctx.String(http.StatusBadRequest, "Usuário não encontrado")
		return
	}

	if !util.CheckPasswordHash(password, user.Password) {
		ctx.String(http.StatusBadRequest, "Senha incorreta")
		return
	}

	ctx.SetCookie("username", user.Username, 3600, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
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