package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/dtos"
	dtosV1 "mobin.dev/internals/dtos/v1"
	"mobin.dev/internals/service"
	"mobin.dev/pkg/constants"
	"mobin.dev/pkg/pagination"
)

type NotesHandler struct {
	s *service.NotesService
}

func NewNotesHandler(s *service.NotesService) *NotesHandler {
	return &NotesHandler{
		s,
	}
}

func (h *NotesHandler) GetNotesHandler(ctx *gin.Context) {
	limit, errPage := strconv.Atoi(ctx.DefaultQuery("limit", strconv.Itoa(constants.DefaultLimit)))
	cursorStr := ctx.DefaultQuery("cursor", "")

	if errPage != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse, dtos.CursorBasedResponseMeta]{
			Success: false,
			Message: "Invalid 'limit' Query Params!",
		})
		return
	}

	if limit > constants.MaxPerPage {
		limit = constants.MaxPerPage
	}

	var cursor pagination.Cursor

	if cursorStr == "" {
		cursor = pagination.Cursor{}
	} else {
		c, err := pagination.DecodeCursor(cursorStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.ApiResponseList[[]dtosV1.NoteResponse, dtos.CursorBasedResponseMeta]{
				Success: false,
				Message: "Invalid Cursor",
			})
			return
		}
		cursor = c
	}

	notes, meta, err := h.s.GetNotes(ctx.Request.Context(), cursor, limit)

	if err != nil {
		status := http.StatusInternalServerError
		message := "Failed to fetch notes"

		if errors.Is(err, service.ErrorInvalidCursor) {
			status = http.StatusBadRequest
			message = "Invalid Cursor"
		}

		ctx.JSON(status, dtos.ApiResponseList[[]dtosV1.NoteResponse, dtos.CursorBasedResponseMeta]{
			Success: false,
			Message: message,
		})
		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponseList[[]dtosV1.NoteResponse, dtos.CursorBasedResponseMeta]{
		Success: true,
		Data:    notes,
		Message: "Get Notes Successfully!",
		Meta:    meta,
	})
}

func (h *NotesHandler) GetNoteHandler(ctx *gin.Context) {
	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
			Success: false,
			Message: "Invalid Note Id",
			Data:    nil,
		})
		return
	}

	note, err := h.s.GetNote(ctx.Request.Context(), id)

	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			ctx.JSON(http.StatusNotFound, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
				Success: false,
				Message: "Note Not Found!",
				Data:    note,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
				Success: false,
				Message: "Failed To Fetch Note!",
				Data:    note,
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, dtos.ApiResponseSingle[dtosV1.NoteResponse]{
		Success: true,
		Data:    note,
		Message: "Get Note Successfully!",
	})
}
func (h *NotesHandler) CreateDummyNotesHandler(ctx *gin.Context) {
	var dummyNotesBody dtosV1.DummyNoteRequest

	if err := ctx.ShouldBindBodyWithJSON(&dummyNotesBody); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[any]{
			Message: "Size must be greater than 0",
			Success: false,
		})
		return
	}

	if err := h.s.CreateDummyNotes(ctx.Request.Context(), dummyNotesBody.Size); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.ApiResponseSingle[any]{
			Message: "Failed to create dummy notes",
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dtos.ApiResponseSingle[dtosV1.DummyNoteRequest]{
		Message: "Dummy Notes Created Successfully",
		Data: &dtosV1.DummyNoteRequest{
			Size: dummyNotesBody.Size,
		},
		Success: true,
	})
}
