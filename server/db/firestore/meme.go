package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/implocell/dashy/core/meme"
)

type memeDB struct {
	db         *firestore.Client
	collection string
}

func NewMemeDB(db *firestore.Client) memeDB {
	return memeDB{
		db:         db,
		collection: "memes",
	}
}

func (db *memeDB) GetByID(ctx context.Context, id string) (*meme.SerializableMeme, error) {
	dnsap, err := db.db.Collection(db.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var meme meme.SerializableMeme
	dnsap.DataTo(&meme)
	return &meme, nil
}