package storage

import "context"

type Storage[T any] interface {
	Upload(ctx context.Context, obj T) error
}
