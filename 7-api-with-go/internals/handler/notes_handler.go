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
	"mobin.dev/pkg/response"
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
		response.Error(ctx, http.StatusBadRequest, "Invalid 'limit' Query Params!")
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
			response.Error(ctx, http.StatusBadRequest, "Invalid Cursor!")
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

		response.Error(ctx, status, message)
		return
	}

	response.ResponseList(ctx, http.StatusOK, true, "Get Notes Successfully!", notes, meta)
}

func (h *NotesHandler) GetNoteHandler(ctx *gin.Context) {
	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid Note Id")
		return
	}

	note, err := h.s.GetNote(ctx.Request.Context(), id)

	if err != nil {
		if errors.Is(err, service.ErrNoteNotFound) {
			response.Error(ctx, http.StatusNotFound, "Note Not Found!")
		} else {
			response.Error(ctx, http.StatusInternalServerError, "Failed To Fetch Note!")
		}

		return
	}

	response.OK(ctx, "Get Note Successfully!", note)
}
func (h *NotesHandler) CreateDummyNotesHandler(ctx *gin.Context) {
	var dummyNotesBody dtosV1.DummyNoteRequest

	if err := ctx.ShouldBindBodyWithJSON(&dummyNotesBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Size must be greater than 0")
		return
	}

	if err := h.s.CreateDummyNotes(ctx.Request.Context(), dummyNotesBody.Size); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to create dummy notes")
		return
	}

	// for this case i cannot may be use that reusable response
	ctx.JSON(http.StatusCreated, dtos.ApiResponseSingle[dtosV1.DummyNoteRequest]{
		Message: "Dummy Notes Created Successfully",
		Data: &dtosV1.DummyNoteRequest{
			Size: dummyNotesBody.Size,
		},
		Success: true,
	})

	// response.OK(ctx, "Dummy Notes Created Successfully")
}
