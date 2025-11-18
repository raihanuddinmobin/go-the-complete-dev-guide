package application

import (
	"errors"

	"mobin.dev/internal/domain/notes/domain"
)

var (
	ErrNotesNotFound = errors.New("Notes not found")
	ErrNoteNotFound  = domain.ErrNoteNotFound
	ErrDBFailure     = domain.ErrDBFailure
	ErrDuplicateNote = domain.ErrDuplicateNote
)
