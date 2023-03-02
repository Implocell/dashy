package storage

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, r io.Reader, fileName string) (string, error)
}
