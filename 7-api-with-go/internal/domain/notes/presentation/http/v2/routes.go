package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterNotesRouter(g *gin.RouterGroup, h *NotesHandler) {

	notes := g.Group("/notes")
	{
		notes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello World FROM v2"})
		})
	}
}
