package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/Registrador-de-Treino/handler"
)

func Routes(router *gin.Engine) {
	// Main page
	router.GET("/", handler.Index)
	router.GET("/hi", handler.Hi)

	// Authentication
	router.GET("/register", handler.RegisterView)
	router.GET("/login", handler.LoginView)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	//router.GET("/validate/email", handler.ValidateEmail)
	//router.GET("/validate/username", handler.ValidateUsername)
	router.POST("/confirmPass", handler.ConfirmPass)
}