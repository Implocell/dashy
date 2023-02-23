package meme

import (
	"context"

	"github.com/implocell/dashy/storage"
)

type MemeStorage interface {
	Upload(ctx context.Context, obj []byte) error
}

type MemeStorageService struct {
	storage storage.Storage[[]byte]
}

func NewMemeStorageService(storage storage.Storage[[]byte]) *MemeStorageService {
	return &MemeStorageService{
		storage: storage,
	}
}

func (s MemeStorageService) Upload(ctx context.Context, obj []byte) error {
	err := s.storage.Upload(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}
