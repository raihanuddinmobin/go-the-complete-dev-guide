package application

import (
	"encoding/xml"
	"time"

	"github.com/go-playground/validator/v10"
	"mobin.dev/internal/domain/notes/domain"
)

type NoteDTO struct {
	XMLName   xml.Name  `xml:"note" json:"-"`
	Id        int64     `json:"id" xml:"id"`
	UserId    int64     `json:"userId" validate:"required"  xml:"userId"`
	Title     string    `json:"title" validate:"required"  xml:"title"`
	Body      string    `json:"body" validate:"required"  xml:"body"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
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

var validate = validator.New()

func (n *NoteDTO) Validate() error {
	return validate.Struct(n)
}
