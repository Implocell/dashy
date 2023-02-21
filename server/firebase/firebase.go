package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type FirebaseContext struct {
	App       *firebase.App
	Firestore *firestore.Client
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

	return &FirebaseContext{
		App:       app,
		Firestore: firestore,
	}, nil
}
