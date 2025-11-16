package application

import (
	"context"
	"errors"
	"fmt"

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

func (s *NotesService) FetchNote(ctx context.Context, id int) (*NoteDTO, error) {
	note, err := s.repo.FindById(ctx, id)

	fmt.Print(err)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			return nil, ErrNoteNotFund
		}
		return nil, ErrDBFailure
	}

	convertedNote := ToNoteDto(note)

	return convertedNote, nil
}
