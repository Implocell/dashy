package firebase

import (
	"context"
	"log"

	"firebase.google.com/go/storage"
)

type memeStorage struct {
	storage *storage.Client
}

func NewMemeStorage(storage *storage.Client) memeStorage {
	return memeStorage{
		storage: storage,
	}
}

func (s memeStorage) Upload(ctx context.Context, obj []byte) error {
	handler, err := s.storage.DefaultBucket()
	if err != nil {
		return err
	}

	writer := handler.Object("something").NewWriter(ctx)
	n, err := writer.Write(obj)
	if err != nil {
		return err
	}

	log.Printf("written %d bytes", n)

	return nil
}
