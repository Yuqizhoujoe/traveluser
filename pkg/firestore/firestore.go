package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

func NewClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, "projectID")
}
