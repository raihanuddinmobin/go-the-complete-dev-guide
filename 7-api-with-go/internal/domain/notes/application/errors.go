package application

import (
	"errors"

	"mobin.dev/internal/domain/notes/domain"
)

var (
	ErrNotesNotFound = errors.New("Notes not found")
	ErrNoteNotFund   = domain.ErrNoteNotFound
	ErrDBFailure     = errors.New("Database failure")
)
