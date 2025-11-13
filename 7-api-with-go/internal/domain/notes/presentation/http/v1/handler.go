package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internal/domain/notes/application"
)

type NotesHandler struct {
	service *application.NoteService
}

func NewNoteHandler(service *application.NoteService) *NotesHandler {
	return &NotesHandler{service}
}

func (n *NotesHandler) GetNotesHandler(ctx *gin.Context) {

	notes, err := n.service.GetNotes(ctx.Request.Context())

	fmt.Println(notes, err)

	ctx.JSON(http.StatusOK, gin.H{"message": "Get All Notes Successfully!"})
}
