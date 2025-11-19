package domain

import "errors"

var (
	ErrNoteNotFound  = errors.New("note not found")
	ErrDuplicateNote = errors.New("duplicate note")
)
