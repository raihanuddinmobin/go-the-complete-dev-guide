package application

import (
	"context"

	"mobin.dev/internal/domain/notes/domain"
)

type NotesService struct {
	repo domain.Repository
}

func NewNotesService(repo domain.Repository) *NotesService {
	return &NotesService{repo}
}

func (s *NotesService) FetchNotes(ctx context.Context) ([]*NoteDTO, error) {
	notes, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, ErrDBFailure
	}

	convertedNots := ToNoteDtos(notes)

	if len(convertedNots) == 0 {

		return nil, ErrNotesNotFound
	}

	return convertedNots, nil
}
