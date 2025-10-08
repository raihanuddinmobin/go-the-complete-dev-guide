package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"mobin.dev/internals/dtos"
	dtosV1 "mobin.dev/internals/dtos/v1"
	"mobin.dev/internals/repository"
	"mobin.dev/pkg/pagination"
)

var (
	ErrNoteNotFound       = errors.New("note not found")
	ErrNotesFetchFailed   = errors.New("failed to fetch notes")
	ErrorDummyNotesInsert = errors.New("failed to insert dummy notes")
	ErrorInvalidCursor    = errors.New("invalid cursor")
)

type NotesService struct {
	r *repository.NotesRepo
}

func NewNotesService(r *repository.NotesRepo) *NotesService {
	return &NotesService{
		r,
	}
}

func (s *NotesService) GetNotes(ctx context.Context, cursor pagination.Cursor, limit int) ([]dtosV1.NoteResponse, *dtos.CursorBasedResponseMeta, error) {
	notes, hasNext, err := s.r.FetchNotes(ctx, cursor, limit)

	if err != nil {
		return nil, &dtos.CursorBasedResponseMeta{}, fmt.Errorf("%w : %v", ErrNotesFetchFailed, err)
	}

	var nextCursor string

	if len(notes) == limit {
		last := notes[len(notes)-1]
		newCursor := pagination.Cursor{CreateAt: last.CreatedAt, Id: last.Id}
		nextCursor, err = pagination.EncodeCursor(newCursor)

		if err != nil {
			return nil, &dtos.CursorBasedResponseMeta{}, fmt.Errorf("%w : %v", ErrorInvalidCursor, err)
		}

	}

	meta := &dtos.CursorBasedResponseMeta{
		Limit:         limit,
		HasNext:       hasNext,
		NextCursor:    nextCursor,
		ReturnedCount: len(notes),
	}

	return dtosV1.ToNoteResponses(notes), meta, nil
}

func (s *NotesService) GetNote(ctx context.Context, id int) (*dtosV1.NoteResponse, error) {

	note, err := s.r.FetchNote(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}

	noteRes := dtosV1.ToNoteResponse(*note)

	return &noteRes, nil
}
func (s *NotesService) CreateDummyNotes(ctx context.Context, size int) error {

	err := s.r.PostDummyNotes(ctx, size)

	if err != nil {
		return fmt.Errorf("create dummy notes failed: %w", err)
	}

	return nil
}
