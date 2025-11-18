package domain

import "errors"

var (
	ErrNoteNotFound  = errors.New("Note not found")
	ErrDuplicateNote = errors.New("Duplicate Creating New Note")
	ErrDBFailure     = errors.New("Database failure")
)
