package v1

import (
	"github.com/gin-gonic/gin"
	"mobin.dev/internal/domain/notes/application"
)

type NotesHandler struct {
	service *application.NotesService
}

func NewNotesHandler(service *application.NotesService) *NotesHandler {
	return &NotesHandler{service}
}

func (h *NotesHandler) GetNotesHandler(ctx *gin.Context) {

}
