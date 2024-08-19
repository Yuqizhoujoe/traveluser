package db

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func NewClient(ctx context.Context) (*firestore.Client, error) {
	// get the project ID from env
	projectId := os.Getenv("FIREBASE_PROJECT_ID")
	if projectId == "" {
		log.Fatalf("FIREBASE_PROJECT_ID is not set in the env variables")
		// err := fmt.Errorf("FIREBASE_PROJECT_ID is not set in the env variables")
		err := errors.New("Firebase project ID is not setup!")
		return nil, err
	}

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
		return nil, err
	}

	return client, nil
}
