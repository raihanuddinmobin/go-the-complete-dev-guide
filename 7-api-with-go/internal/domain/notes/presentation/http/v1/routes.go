package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterNotesRouter(g *gin.RouterGroup, h *NotesHandler) {

	notes := g.Group("/notes")
	{
		notes.GET("/", h.GetNotesHandler)
	}
}
