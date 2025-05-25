package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/Registrador-de-Treino/handler"
	"github.com/mickaelyoshua/Registrador-de-Treino/middleware"
)

func Routes(router *gin.Engine) {
	// Middleware
	authenticate := router.Group("/")
	authenticate.Use(middleware.Authenticate)

	// Main page
	authenticate.GET("/", handler.Index)
	authenticate.GET("/hi", handler.Hi)

	// Workout
	authenticate.GET("/workout", handler.WorkoutView)
	authenticate.GET("/workout/create", handler.WorkoutCreateView)
	authenticate.POST("/workout/create", handler.WorkoutCreate)
	authenticate.DELETE("/workout/delete/:id", handler.WorkoutDelete)

	// Authentication
	router.GET("/register", handler.RegisterView)
	router.GET("/login", handler.LoginView)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	//router.GET("/validate/email", handler.ValidateEmail)
	//router.GET("/validate/username", handler.ValidateUsername)
	router.POST("/confirmPass", handler.ConfirmPass)
}