package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"mobin.dev/internals/model"
	"mobin.dev/pkg/pagination"
)

type NotesRepo struct {
	db *sql.DB
}

func NewNotesRepo(db *sql.DB) *NotesRepo {
	return &NotesRepo{
		db,
	}
}

func (r *NotesRepo) FetchNotes(ctx context.Context, cursor pagination.Cursor, limit int) ([]model.Note, bool, error) {
	baseQuery := `SELECT id, user_id, title, body, tags, created_at, updated_at FROM notes `

	fetchLimit := limit + 1
	var args []interface{}
	var whereClause string

	query := baseQuery

	if !cursor.IsEmpty() {
		whereClause = `WHERE (created_at, id) < ($1, $2)`
		args = append(args, cursor.CreateAt, cursor.Id)
	}
	if whereClause != "" {
		query += whereClause + " ORDER BY created_at DESC, id DESC LIMIT $3"
	} else {
		query += ` ORDER BY created_at DESC, id DESC LIMIT $1`
	}

	args = append(args, fetchLimit)

	fmt.Println(query, args)
	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, false, fmt.Errorf("failed to query db data %w ", err)
	}

	defer rows.Close()

	notes := make([]model.Note, 0, limit)

	for rows.Next() {
		var note model.Note
		var tags json.RawMessage

		if err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, false, fmt.Errorf("failed to scan rows : %w", err)
		}

		err := json.Unmarshal(tags, &note.Tags)
		if err != nil {
			return nil, false, fmt.Errorf("error unmarshalling JSON : %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, false, fmt.Errorf("row iteration error: %w", err)
	}

	if len(notes) == 0 {
		return notes, false, nil
	}

	hasNext := false
	if len(notes) > limit {
		hasNext = true
		notes = notes[:limit]
	}

	return notes, hasNext, nil
}

func (r *NotesRepo) FetchNote(ctx context.Context, id int) (*model.Note, error) {
	fmt.Println(id)
	row := r.db.QueryRowContext(ctx, `SELECT id, user_id, title, body, tags, created_at, updated_at
					FROM notes WHERE id = $1`, id)

	var note model.Note
	var tags json.RawMessage

	err := row.Scan(&note.Id, &note.UserId, &note.Title, &note.Body, &tags, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("failed to scan note: %w", err)

	}

	err = json.Unmarshal(tags, &note.Tags)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON : %w", err)
	}

	return &note, nil
}

func (r NotesRepo) PostDummyNotes(ctx context.Context, size int) error {
	tagJson, _ := json.Marshal(map[string]string{
		gofakeit.Sentence(5): gofakeit.Sentence(10),
	})

	for range size {

		min := 31
		max := 45
		randomNumber := rand.Intn(max-min+1) + min
		note := model.Note{
			UserId:    randomNumber,
			Title:     gofakeit.Sentence(10),
			Body:      gofakeit.Paragraph(3, 5, 30, " "),
			CreatedAt: gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
			UpdatedAt: gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
			Tags:      tagJson,
		}

		_, err := r.db.ExecContext(ctx, `INSERT INTO notes (user_id, title, body, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
			note.UserId,
			note.Title,
			note.Body,
			note.Tags,
			note.CreatedAt,
			note.UpdatedAt)

		if err != nil {
			return fmt.Errorf("failed to insert dummy notes : %w ", err)
		}
	}

	return nil
}
