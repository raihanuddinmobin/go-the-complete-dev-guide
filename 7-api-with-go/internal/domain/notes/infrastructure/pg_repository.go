package infrastructure

import (
	"context"
	"database/sql"
	"strings"

	"go.uber.org/zap"
	"mobin.dev/internal/domain/notes/domain"
	"mobin.dev/internal/infrastructure/logger"
)

type NotesRepository struct {
	db *sql.DB
}

func NewNotesRepository(db *sql.DB) domain.Repository {
	return &NotesRepository{db}
}

func (r *NotesRepository) Create(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO notes (user_id, title, body) VALUES($1, $2, $3) RETURNING id, created_at, updated_at`,
		note.UserId, note.Title, note.Body,
	).Scan(&note.Id, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		logger.Log.Error("DB Error on Create()", zap.Error(err))

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, domain.ErrDuplicateNote
		}
		return nil, err
	}

	return note, nil
}

func (r *NotesRepository) FindAll(ctx context.Context) ([]*domain.Note, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes LIMIT 200`)

	if err != nil {
		logger.Log.Error("DB Error on FindAll()", zap.Error(err))
		return nil, nil
	}
	defer rows.Close()

	var notes []*domain.Note

	for rows.Next() {
		var n domain.Note

		if err := rows.Scan(&n.Id, &n.UserId, &n.Title, &n.Body, &n.CreatedAt, &n.UpdatedAt); err != nil {
			logger.Log.Error("Scan Error in FindAll()", zap.Int64("id", n.Id), zap.String("title", n.Title), zap.Error(err))
			return nil, err
		}

		notes = append(notes, &n)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Error("Rows Iteration Error in FindAll()", zap.Error(err))
		return nil, err
	}

	return notes, nil
}

func (r *NotesRepository) FindById(ctx context.Context, id int64) (*domain.Note, error) {
	var note = &domain.Note{}

	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes WHERE id = $1`, id)
	if err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt); err != nil {
		logger.Log.Error("DB Error in FindById()", zap.Error(err))

		if err == sql.ErrNoRows {
			return nil, domain.ErrNoteNotFound
		}
		return nil, err
	}

	return note, nil
}
