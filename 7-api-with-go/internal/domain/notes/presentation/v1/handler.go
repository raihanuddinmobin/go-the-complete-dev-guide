package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mobin.dev/internal/common/errcode"
	"mobin.dev/internal/common/response"
	"mobin.dev/internal/common/validation"
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
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", errcode.INTERNAL_SERVER_ERROR)
		}

		return
	}

	response.Success(c, "Get all notes successfully!", notes)
}

func (h *NotesHandler) GetNoteHandler(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)

	if id == "" || err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Path parameters", errcode.INTERNAL_SERVER_ERROR, gin.H{
			"message": "Path parameters must be number or greater then 1",
			"err":     id,
		})
		return
	}

	note, err := h.service.FetchNote(c, num)

	if err != nil {
		switch {
		case errors.Is(err, application.ErrNotesNotFound):
			response.Error(c, http.StatusNotFound, "Note Not Found", errcode.NOT_FOUND)
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", errcode.INTERNAL_SERVER_ERROR)
		}
		return
	}

	response.Success(c, "Fetch Notes Successfully!", note)
}

func (h *NotesHandler) PostNoteHandler(c *gin.Context) {

	err := errors.New("test")

	var dto application.NoteDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Invalid JSON Body Data!", errcode.INTERNAL_SERVER_ERROR)
		return
	}

	if err := dto.Validate(); err != nil {
		if vErrs, ok := err.(validator.ValidationErrors); ok {
			response.Error(
				c,
				http.StatusBadRequest,
				"Validation Failed",
				errcode.VALIDATION_ERROR,
				validation.FormatValidationErrors(vErrs),
			)
		}

		return
	}

	createdNote, err := h.service.PostNote(c, &dto)

	if err != nil {
		switch {
		case errors.Is(err, application.ErrDuplicateNote):
			response.Error(c, http.StatusConflict, "Duplicating Note", errcode.DUPLICATE)
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", errcode.INTERNAL_SERVER_ERROR)
		}
		return
	}

	response.Success(c, "Notes Created Successfully", createdNote, http.StatusCreated)
}
