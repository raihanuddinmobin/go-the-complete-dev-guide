package application

import (
	"fmt"
	"time"

	"mobin.dev/internal/domain/notes/domain"
)

type NoteDTO struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToNoteDTO(note *domain.Note) *NoteDTO {
	return &NoteDTO{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	}
}

func ToNoteDTOs(notes []*domain.Note) []*NoteDTO {
	dtos := make([]*NoteDTO, 0, len(notes))

	for _, note := range notes {
		dtos = append(dtos, ToNoteDTO(note))
	}

	fmt.Println(dtos)

	return dtos
}
