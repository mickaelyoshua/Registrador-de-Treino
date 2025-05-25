package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/mickaelyoshua/Registrador-de-Treino/model"
	"github.com/mickaelyoshua/Registrador-de-Treino/util"
	"github.com/mickaelyoshua/Registrador-de-Treino/view"
)

func Render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}
func HandleRenderError(err error) {
	if err != nil {
		log.Printf("Could not render template: %v", err)
	}
}

// Main page
func Index(ctx *gin.Context) {
	retrievedToken, err := util.GetTokenFromCookie(ctx)
	if err != nil {
		log.Printf("Error getting token from cookie: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user, err := model.FindUserByToken(retrievedToken)
	if err != nil {
		log.Printf("Error finding user: \n%v", err)
		return
	}

	err = Render(ctx, http.StatusOK, view.Index(user))
	HandleRenderError(err)
}
func Hi(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		log.Printf("Error getting cookie: \n%v", err)
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
	username := ctx.Request.FormValue("username")
	email := ctx.Request.FormValue("email")
	password, err := util.HashPassword(ctx.Request.FormValue("password"))
	if err != nil {
		log.Printf("Error hashing password: \n%v", err)
	}

	created := time.Now()
	user := model.NewUser(username, email, password, created, created)
	err = user.Save()
	if err != nil {
		log.Printf("Error saving User: \n%v", err)
		return
	}

	token, err := util.GenerateToken(email, user.Id) // Pass ObjectID directly
	if err != nil {
		log.Printf("Error generating token: \n%v", err)
		ctx.String(http.StatusBadRequest, "Erro ao gerar token")
		return
	}
	ctx.SetCookie("token", token, int(2*time.Hour.Seconds()), "/", "", false, true)

	ctx.Redirect(http.StatusSeeOther, "/")
}

func Login(ctx *gin.Context) {
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	user, err := model.FindUserByFilter(bson.M{"email": email})
	if err != nil {
		log.Printf("Error finding user by email: \n%v", err)
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

	token, err := util.GenerateToken(email, user.Id) // Pass ObjectID directly
	if err != nil {
		log.Printf("Error generating token: \n%v", err)
		ctx.String(http.StatusBadRequest, "Erro ao gerar token")
		return
	}
	ctx.SetCookie("token", token, int(2*time.Hour.Seconds()), "/", "", false, true)

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

// Workout
func WorkoutView(ctx *gin.Context) {
	retrievedToken, err := util.GetTokenFromCookie(ctx)
	if err != nil {
		log.Printf("Error getting token from cookie: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user, err := model.FindUserByToken(retrievedToken)
	if err != nil {
		log.Printf("Error finding user: \n%v", err)
		return
	}

	workouts, err := model.FindAllWorkoutsByUserId(user.Id)
	if err != nil {
		log.Printf("Error finding workouts: \n%v", err)
		return
	}

	err = Render(ctx, http.StatusOK, view.Workouts(workouts))
	HandleRenderError(err)
}

func WorkoutCreateView(ctx *gin.Context) {
	retrievedToken, err := util.GetTokenFromCookie(ctx)
	if err != nil {
		log.Printf("Error getting token from cookie: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	_, err = model.FindUserByToken(retrievedToken)
	if err != nil {
		log.Printf("Error finding user: \n%v", err)
		return
	}

	err = Render(ctx, http.StatusOK, view.WorkoutCreate())
	HandleRenderError(err)
}

func WorkoutCreate(ctx *gin.Context) {
	retrievedToken, err := util.GetTokenFromCookie(ctx)
	if err != nil {
		log.Printf("Error getting token from cookie: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user, err := model.FindUserByToken(retrievedToken)
	if err != nil {
		log.Printf("Error finding user: \n%v", err)
		return
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	created := time.Now()

	workout := model.NewWorkout(title, description, user.Id, created, created)
	err = workout.Save()
	if err != nil {
		log.Printf("Error saving workout: \n%v", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/workout")
}

func WorkoutDelete(ctx *gin.Context) {
	retrievedToken, err := util.GetTokenFromCookie(ctx)
	if err != nil {
		log.Printf("Error getting token from cookie: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	_, err = model.FindUserByToken(retrievedToken)
	if err != nil {
		log.Printf("Error finding user: \n%v", err)
		return
	}

	workoutId := ctx.Param("id")
	workout, err := model.GetWorkoutById(workoutId)
	if err != nil {
		log.Printf("Error getting workout: \n%v", err)
		return
	}
	err = workout.Delete()
	if err != nil {
		log.Printf("Error deleting workout: \n%v", err)
		return
	}

	ctx.Status(http.StatusOK)
}