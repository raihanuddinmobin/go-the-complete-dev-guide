package domain

import "context"

type Repository interface {
	Create(ctx context.Context, note *Note) error
}
