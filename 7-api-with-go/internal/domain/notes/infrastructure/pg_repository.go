package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *NotesRepository) FindAll(ctx context.Context) ([]*domain.Note, error) {

	rows, err := r.db.QueryContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes LIMIT 200`)

	if err != nil {
		return nil, fmt.Errorf("query error in FetchAll : %w ", err)
	}
	defer rows.Close()

	var notes []*domain.Note

	for rows.Next() {
		var n domain.Note

		if err := rows.Scan(&n.Id, &n.UserId, &n.Title, &n.Body, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan error in FetchAll : %w ", err)
		}

		notes = append(notes, &n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error in FetchAll : %w ", err)
	}

	return notes, nil
}

func (r *NotesRepository) FindById(ctx context.Context, id int) (*domain.Note, error) {
	var note = &domain.Note{}

	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes WHERE id = $1`, id)
	if err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoteNotFound
		}
		return nil, fmt.Errorf("Failed to fetch notes Id : %d  , reason : %w ", id, err)
	}

	return note, nil
}
