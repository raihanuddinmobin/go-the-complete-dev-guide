package application

import "mobin.dev/internal/domain/notes/domain"

type NotesService struct {
	repo domain.Repository
}

func NewNotesService(repo domain.Repository) *NotesService {
	return &NotesService{repo}
}
