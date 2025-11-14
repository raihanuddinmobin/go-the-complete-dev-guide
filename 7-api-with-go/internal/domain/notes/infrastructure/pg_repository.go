package infrastructure

import (
	"context"
	"database/sql"

	"mobin.dev/internal/domain/notes/domain"
)

type NotesRepository struct {
	db *sql.DB
}

func NewNotesRepository(db *sql.DB) domain.Repository {
	return &NotesRepository{db}
}

func (r *NotesRepository) Create(ctx context.Context, note *domain.Note) error {
	return nil
}
