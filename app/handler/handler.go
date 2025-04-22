package handler

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
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

func Index(ctx *gin.Context) {
	err := Render(ctx, http.StatusOK, view.Index())
	HandleRenderError(err)
}