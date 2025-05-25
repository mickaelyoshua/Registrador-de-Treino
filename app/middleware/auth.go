package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/mickaelyoshua/Registrador-de-Treino/util"
	"github.com/mickaelyoshua/Registrador-de-Treino/model"
)

func Authenticate(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	retrievedToken, err := util.ValidateToken(token)
	if err != nil {
		log.Printf("Error validating token: \n%v", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Check if the token is expired
	if int64(retrievedToken["expiration"].(float64)) < time.Now().Unix() {
		log.Printf("Token expired: %v", retrievedToken["expiration"])
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	email := retrievedToken["email"].(string)
	user, err := model.FindUserByFilter(bson.M{"email": email})
	if err != nil {
		log.Printf("Error finding user by email: \n%v", err)
		ctx.String(http.StatusBadRequest, "Usuário não encontrado")
		return
	}

	token, err = util.GenerateToken(email, user.Id) // Pass ObjectID directly
	if err != nil {
		log.Printf("Error generating token: \n%v", err)
		ctx.String(http.StatusBadRequest, "Erro ao gerar token")
		return
	}
	ctx.SetCookie("token", token, int(2*time.Hour.Seconds()), "/", "", false, true)
}