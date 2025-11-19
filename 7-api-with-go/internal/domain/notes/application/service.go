package application

import (
	"context"
	"errors"

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
		return nil, ErrNoteNotFound
	}

	return convertedNots, nil
}

func (s *NotesService) FetchNote(ctx context.Context, id int64) (*NoteDTO, error) {
	note, err := s.repo.FindById(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) {
			return nil, ErrNoteNotFound
		}
		return nil, ErrDBFailure
	}

	return ToNoteDto(note), nil
}

func (s *NotesService) PostNote(ctx context.Context, nDto *NoteDTO) (*NoteDTO, error) {
	note := domain.NewNote(nDto.UserId, nDto.Title, nDto.Body)

	createdNote, err := s.repo.Create(ctx, note)

	if err != nil {
		if errors.Is(err, domain.ErrDuplicateNote) {
			return nil, ErrDuplicateNote
		}
		return nil, ErrDBFailure
	}

	return ToNoteDto(createdNote), nil
}
