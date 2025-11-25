package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mobin.dev/internal/common/errcode"
	"mobin.dev/internal/common/pagination"
	"mobin.dev/internal/common/response"
	"mobin.dev/internal/common/validation"
	"mobin.dev/internal/domain/notes/application"
	"mobin.dev/internal/infrastructure/logger"
)

type NotesHandler struct {
	service *application.NotesService
}

func NewNotesHandler(service *application.NotesService) *NotesHandler {
	return &NotesHandler{service}
}

func (h *NotesHandler) GetNotesHandler(c *gin.Context) {
	traceId := logger.TraceIdFromContext(c.Request.Context())

	limit, lErr := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, oErr := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if lErr != nil || oErr != nil {
		response.Error(c, http.StatusBadRequest, "Limit, Offset Must Be An Integer", traceId, errcode.FIELD_REQUIRED, gin.H{
			"message": "You Provided Limit , Offset",
			"limit":   c.Query("limit"),
			"offset":  c.Query("offset"),
		})
		return
	} else if limit < 0 || offset < 0 {
		response.Error(c, http.StatusBadRequest, "Limit, Offset Must Be Positive", traceId, errcode.FIELD_REQUIRED, gin.H{
			"message": "You Provided Limit , Offset",
			"limit":   limit,
			"offset":  offset,
		})
		return
	}

	if limit > pagination.MAX_LIMIT {
		limit = pagination.MAX_LIMIT
	} else if limit < pagination.MIN_LIMIT {
		limit = pagination.MIN_LIMIT
	}

	fmt.Println(limit, offset)

	notes, err := h.service.FetchNotes(c.Request.Context(), limit, offset)

	if err != nil {
		switch {
		case errors.Is(err, application.ErrNoteNotFound):
			response.Error(c, http.StatusNotFound, "No Notes Found", traceId, errcode.NOT_FOUND)
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", traceId, errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", traceId, errcode.INTERNAL_SERVER_ERROR)
		}

		return
	}

	response.Success(c, "Get all notes successfully!", notes)
}

func (h *NotesHandler) GetNoteHandler(c *gin.Context) {
	traceId := logger.TraceIdFromContext(c.Request.Context())

	id := c.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid Path parameters", traceId, errcode.INTERNAL_SERVER_ERROR, gin.H{
			"message": "Path parameters must be number or greater then 1",
			"err":     id,
		})
		return
	}

	note, err := h.service.FetchNote(c.Request.Context(), num)

	if err != nil {
		switch {
		case errors.Is(err, application.ErrNoteNotFound):
			response.Error(c, http.StatusNotFound, "Note Not Found", traceId, errcode.NOT_FOUND)
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", traceId, errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", traceId, errcode.INTERNAL_SERVER_ERROR)
		}
		return
	}

	response.Success(c, "Fetch Notes Successfully!", note)
}

func (h *NotesHandler) PostNoteHandler(c *gin.Context) {
	traceId := logger.TraceIdFromContext(c.Request.Context())

	var dto application.NoteDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Invalid JSON Body Data!", traceId, errcode.INTERNAL_SERVER_ERROR)
		return
	}

	if err := dto.Validate(); err != nil {
		if vErrs, ok := err.(validator.ValidationErrors); ok {
			response.Error(
				c,
				http.StatusBadRequest,
				"Validation Failed",
				traceId,
				errcode.VALIDATION_ERROR,
				validation.FormatValidationErrors(vErrs),
			)
		}

		return
	}

	createdNote, err := h.service.PostNote(c.Request.Context(), &dto)

	if err != nil {
		switch {
		case errors.Is(err, application.ErrDuplicateNote):
			response.Error(c, http.StatusConflict, "Duplicating Note", traceId, errcode.DUPLICATE)
			return
		case errors.Is(err, application.ErrDBFailure):
			response.Error(c, http.StatusInternalServerError, "Unexpected server error", traceId, errcode.INTERNAL_SERVER_ERROR)
			return
		default:
			response.Error(c, http.StatusInternalServerError, "Something went wrong", traceId, errcode.INTERNAL_SERVER_ERROR)
		}
		return
	}

	response.Success(c, "Notes Created Successfully", createdNote, http.StatusCreated)
}
