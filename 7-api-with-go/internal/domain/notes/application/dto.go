package application

import "mobin.dev/internal/domain/notes/domain"

type NoteDTO struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	Title     string `json:"title" validate:"required,min=5,max=100"`
	Body      string `json:"description" validate:"required,min=5,max=1000"`
	CreatedAt string `json:"createdAt" validate:"required"`
	UpdatedAt string `json:"updatedAt" validate:"required"`
}

func ToNoteDto(note *domain.Note) *NoteDTO {
	return &NoteDTO{Id: note.Id, UserId: note.UserId, Title: note.Title, Body: note.Body, CreatedAt: note.CreatedAt, UpdatedAt: note.UpdatedAt}
}

func ToNoteDtos(notes []*domain.Note) []*NoteDTO {
	dtos := make([]*NoteDTO, 0, len(notes))

	for _, note := range notes {
		dtos = append(dtos, ToNoteDto(note))
	}

	return dtos
}
