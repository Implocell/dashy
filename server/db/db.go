package db

import (
	"context"
)

type DB[T any] interface {
	GetByID(ctx context.Context, id string) (T, error)
}
