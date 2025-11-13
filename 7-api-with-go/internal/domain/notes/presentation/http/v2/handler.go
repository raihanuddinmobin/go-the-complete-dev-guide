package v2

import "mobin.dev/internal/domain/notes/application"

type NotesHandler struct {
	service *application.NoteService
}

func NewNoteHandler(service *application.NoteService) *NotesHandler {
	return &NotesHandler{service}
}
