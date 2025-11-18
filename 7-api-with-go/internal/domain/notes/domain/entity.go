package domain

type Note struct {
	Id        int64
	UserId    int64
	Title     string
	Body      string
	CreatedAt string
	UpdatedAt string
}

func NewNote(userId int64, title, body string) *Note {
	return &Note{
		UserId: userId,
		Title:  title,
		Body:   body,
	}
}
