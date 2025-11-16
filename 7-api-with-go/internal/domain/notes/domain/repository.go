package domain

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]*Note, error)
	Create(ctx context.Context, note *Note) error
}
