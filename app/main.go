package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/handlers"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("views/*.html")

	server.GET("/register", handlers.Register)
	server.GET("/login", handlers.Login)
	server.GET("/", handlers.DefaultHandler)

	server.Run(":8080")
}