package domain

import "context"

type Repository interface {
	Save(context context.Context, n *Note) error
	FindAll(context context.Context) ([]*Note, error)
	FindById(context context.Context, id int64) (*Note, error)
	Delete(context context.Context, id int64) error
}
