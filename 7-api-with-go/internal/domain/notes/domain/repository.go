package domain

import "context"

type Repository interface {
	FindAll(ctx context.Context, limit, offset int) ([]*Note, error)
	FindById(ctx context.Context, id int64) (*Note, error)
	Create(ctx context.Context, note *Note) (*Note, error)
}
