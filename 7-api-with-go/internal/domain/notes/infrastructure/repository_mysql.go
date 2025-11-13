package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"mobin.dev/internal/domain/notes/domain"
)

type NotesRepo struct {
	db *sql.DB
}

func NewNoteRepo(db *sql.DB) domain.Repository {
	return &NotesRepo{db}
}

func (r *NotesRepo) Save(context context.Context, n *domain.Note) error {
	return nil
}

func (r *NotesRepo) FindAll(context context.Context) ([]*domain.Note, error) {

	note := &domain.Note{
		ID:          1,
		Title:       "test",
		Description: "test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	notes := []*domain.Note{note, note, note, note}

	return notes, nil
}

func (r *NotesRepo) FindById(context context.Context, id int64) (*domain.Note, error) {
	return nil, nil
}

func (r *NotesRepo) Delete(context context.Context, id int64) error {
	return nil
}
