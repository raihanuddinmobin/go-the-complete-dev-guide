package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterNotesRouter(grp *gin.RouterGroup, h *NotesHandler) {

	notes := grp.Group("/notes")
	{
		notes.GET("/", h.GetNotesHandler)
		notes.GET("/:id", h.GetNoteHandler)
	}

}
