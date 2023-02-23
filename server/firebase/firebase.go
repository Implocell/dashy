package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"

	"google.golang.org/api/option"
)

type FirebaseContext struct {
	App       *firebase.App
	Firestore *firestore.Client
	Storage   *storage.Client
}

func InitAndReturnApp() (*FirebaseContext, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	firestore, err := app.Firestore(ctx)

	if err != nil {
		return nil, fmt.Errorf("error initializing firestore: %v", err)
	}

	storage, err := app.Storage(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to initalize firebase storage: %v", err)
	}

	return &FirebaseContext{
		App:       app,
		Firestore: firestore,
		Storage:   storage,
	}, nil
}
