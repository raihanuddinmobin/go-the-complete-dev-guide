package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internal/common/errcode"
	"mobin.dev/internal/common/response"
	"mobin.dev/internal/domain/notes/application"
)

type NotesHandler struct {
	service *application.NotesService
}

func NewNotesHandler(service *application.NotesService) *NotesHandler {
	return &NotesHandler{service}
}

func (h *NotesHandler) GetNotesHandler(c *gin.Context) {
	notes, err := h.service.FetchNotes(c.Request.Context())

	if err != nil {
		switch {
		case errors.Is(err, application.ErrNotesNotFound):
			response.Error(c, http.StatusNotFound, "No Notes Found", errcode.NOT_FOUND)
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", errcode.INTERNAL_SERVER_ERROR)
		case errors.Is(err, application.ErrNotesNotFound):
			response.Error(c, http.StatusInternalServerError, "Something went wrong", errcode.INTERNAL_SERVER_ERROR)
		}

		return
	}

	response.Success(c, "Get all notes successfully!", notes)
}
