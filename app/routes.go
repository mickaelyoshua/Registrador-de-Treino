package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/handler"
)

func Routes(router *gin.Engine) {
	// Main page
	router.GET("/", handler.Index)

	// Authentication
	router.GET("/register", handler.RegisterView)
	router.POST("/register", handler.Register)
	router.POST("/confirmPass", handler.ConfirmPass)
}