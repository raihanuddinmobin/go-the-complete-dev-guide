package application

import "errors"

var ErrNotesNotFound = errors.New("Notes not found")
var ErrDBFailure = errors.New("Database failure")
