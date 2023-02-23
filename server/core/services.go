package core

import (
	"github.com/implocell/dashy/core/meme"
	"github.com/implocell/dashy/db/firestore"
	"github.com/implocell/dashy/firebase"
	firebaseStorage "github.com/implocell/dashy/storage/firebase"
)

type Services struct {
	memeService *meme.MemeService
}

func SetupServices(firebase *firebase.FirebaseContext) *Services {
	memeDB := firestore.NewMemeDB(firebase.Firestore)
	memeStorage := firebaseStorage.NewMemeStorage(firebase.Storage)
	return &Services{
		memeService: meme.NewMemeService(&memeDB, &memeStorage),
	}
}

func (s *Services) GetMemeService() *meme.MemeService {
	return s.memeService
}
