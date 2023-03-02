package db

import (
	"context"
)

type DB[T any] interface {
	GetByID(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, item *T) error
	GetAll(ctx context.Context) ([]T, error)
}
