package domain

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]*Note, error)
	FindById(ctx context.Context, id int) (*Note, error)
	Create(ctx context.Context, note *Note) (*Note, error)
}
