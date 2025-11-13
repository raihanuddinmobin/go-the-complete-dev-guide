package application

import (
	"context"

	"mobin.dev/internal/domain/notes/domain"
)

type NoteService struct {
	repo domain.Repository
}

func NewNotesService(repo domain.Repository) *NoteService {
	return &NoteService{repo}
}

func (n *NoteService) GetNotes(context context.Context) ([]*domain.Note, error) {
	notes, err := n.repo.FindAll(context)

	return notes, err
}
