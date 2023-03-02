package firebase

import (
	"context"
	"fmt"
	"io"
	"os"

	gStorage "cloud.google.com/go/storage"
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

func (s memeStorage) Upload(ctx context.Context, r io.Reader, fileName string) (string, error) {
	handler, err := s.storage.Bucket(os.Getenv("FIREBASE_BUCKET_MEME"))
	if err != nil {
		return "", err
	}

	o := handler.Object(fmt.Sprintf("memes/%s", fileName))
	writer := o.NewWriter(ctx)

	_, err = io.Copy(writer, r)
	if err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	if err := o.ACL().Set(ctx, gStorage.AllUsers, gStorage.RoleReader); err != nil {
		return "", err
	}

	previewLink := "https://firebasestorage.googleapis.com/v0/b/dashy-9a477.appspot.com/o/memes%2F" + fmt.Sprintf("%s?alt=media", fileName)

	return previewLink, nil
}
