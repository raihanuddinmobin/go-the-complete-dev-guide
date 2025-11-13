package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internal/common/response"
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

	err = errors.New("something went wrong")
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch notes", response.INTERNAL_SERVER_ERROR)
		return
	}

	notesDto := application.ToNoteDTOs(notes)

	response.Success(ctx, "Fetched Notes Successfully", notesDto)
}
