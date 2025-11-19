package application

import (
	"errors"
)

var (
	ErrNoteNotFound  = errors.New("note not found")
	ErrDuplicateNote = errors.New("duplicate note")
	ErrDBFailure     = errors.New("database failure")
)
