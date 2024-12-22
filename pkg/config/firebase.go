package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	App  *firebase.App
	Auth *auth.Client
}

func CreateFirebaseService(cfg Configuration) (*FirebaseService, error) {
	opt := option.WithCredentialsFile(cfg.FirebaseCredentials)

	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase: %w", err)
	}

	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get Firebase Auth client: %w", err)
	}

	return &FirebaseService{
		App:  firebaseApp,
		Auth: firebaseAuth,
	}, nil
}
