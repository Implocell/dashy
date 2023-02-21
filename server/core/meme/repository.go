package meme

import (
	"context"

	"github.com/implocell/dashy/db"
)

type MemeRepository interface {
	GetByID(ctx context.Context, id string) (*SerializableMeme, error)
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
