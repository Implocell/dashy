package meme

import (
	"context"
	"io"

	"github.com/implocell/dashy/storage"
)

type MemeStorage interface {
	Upload(ctx context.Context, r io.Reader, fileName string) (string, error)
}

type MemeStorageService struct {
	storage storage.Storage
}

func NewMemeStorageService(storage storage.Storage) *MemeStorageService {
	return &MemeStorageService{
		storage: storage,
	}
}

func (s MemeStorageService) Upload(ctx context.Context, r io.Reader, fileName string) (string, error) {
	url, err := s.storage.Upload(ctx, r, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}
