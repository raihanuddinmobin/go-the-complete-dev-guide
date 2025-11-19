package domain

import "time"

type Note struct {
	Id        int64
	UserId    int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewNote(userId int64, title, body string) *Note {
	return &Note{
		UserId: userId,
		Title:  title,
		Body:   body,
	}
}
