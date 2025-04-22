package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/Registrador-de-Treino/handler"
)

func Routes(router *gin.Engine) {
	router.GET("/", handler.Index)
}