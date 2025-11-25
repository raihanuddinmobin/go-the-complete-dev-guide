package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"mobin.dev/internal/domain/notes/domain"
	"mobin.dev/internal/infrastructure/logger"
)

type NotesRepository struct {
	db    *sql.DB
	cache *NotesCache
}

func NewNotesRepository(db *sql.DB, redisClient *NotesCache) domain.Repository {
	return &NotesRepository{db, redisClient}
}

func (r *NotesRepository) Create(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	log := logger.LoggerWithContext(ctx)

	err := r.db.QueryRowContext(ctx,
		`INSERT INTO notes (user_id, title, body) VALUES($1, $2, $3) RETURNING id, created_at, updated_at`,
		note.UserId, note.Title, note.Body,
	).Scan(&note.Id, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		log.Error("DB Error on Create()", zap.Error(err))

		if strings.Contains(err.Error(), "duplicate key") {
			return nil, domain.ErrDuplicateNote
		}
		return nil, err
	}

	return note, nil
}

func (r *NotesRepository) FindAll(ctx context.Context) ([]*domain.Note, error) {
	log := logger.LoggerWithContext(ctx)

	rows, err := r.db.QueryContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes LIMIT 200`)

	if err != nil {
		log.Error("DB Error on FindAll()", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var notes []*domain.Note

	for rows.Next() {
		var n domain.Note

		if err := rows.Scan(&n.Id, &n.UserId, &n.Title, &n.Body, &n.CreatedAt, &n.UpdatedAt); err != nil {
			log.Error("Scan Error in FindAll()", zap.Int64("id", n.Id), zap.String("title", n.Title), zap.Error(err))
			return nil, err
		}

		notes = append(notes, &n)
	}

	if err := rows.Err(); err != nil {
		log.Error("Rows Iteration Error in FindAll()", zap.Error(err))
		return nil, err
	}

	return notes, nil
}

func (r *NotesRepository) FindById(ctx context.Context, id int64) (*domain.Note, error) {
	log := logger.LoggerWithContext(ctx)
	cacheNote, err := r.cache.GetNote(ctx, id)

	if err != nil {
		log.Warn("Redis GET Failed", zap.Error(err))
	} else if cacheNote != nil {
		fmt.Println("Cache HIT")
		return cacheNote, nil
	}
	fmt.Println("Cache Miss")

	var note = &domain.Note{}

	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, title, body, created_at, updated_at FROM notes WHERE id = $1`, id)
	if err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt); err != nil {
		log.Error("DB Error in FindById()", zap.Error(err))

		if err == sql.ErrNoRows {
			return nil, domain.ErrNoteNotFound
		}
		return nil, err
	}

	if err := r.cache.SetNote(ctx, id, note); err != nil {
		log.Warn("Redis SET failed", zap.Error(err))
	}

	return note, nil
}
