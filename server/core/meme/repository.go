package meme

import (
	"context"

	"github.com/implocell/dashy/db"
)

type MemeRepository interface {
	GetByID(ctx context.Context, id string) (*SerializableMeme, error)
	GetAll(ctx context.Context) (*[]SerializableMeme, error)
	Create(ctx context.Context, meme *SerializableMeme) error
}

type MemeDatabaseService struct {
	db db.DB[SerializableMeme]
}

func NewMemeDatabaseService(db db.DB[SerializableMeme]) *MemeDatabaseService {
	return &MemeDatabaseService{
		db: db,
	}
}

func (m *MemeDatabaseService) GetByID(ctx context.Context, id string) (*SerializableMeme, error) {
	meme, err := m.db.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &meme, nil
}

func (m *MemeDatabaseService) GetAll(ctx context.Context) (*[]SerializableMeme, error) {
	memes, err := m.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &memes, nil
}

func (m *MemeDatabaseService) Create(ctx context.Context, meme *SerializableMeme) error {
	if err := m.db.Create(ctx, meme); err != nil {
		return err
	}

	return nil
}
