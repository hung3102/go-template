package infrastructure

import (
	"context"
	"fmt"
	"os"
	"time"

	cs "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"

	istorage "gcim/example/internal/domain/storage"
)

type CloudStorage struct {
	client *storage.Client
}

func NewStorage() (istorage.Storage, error) {
	ctx := context.Background()

	app, err := newFirebaseClientForStorage(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewStorage: firebase.NewApp: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewStorage: app.Storage: %v", err)
	}

	return &CloudStorage{
		client: client,
	}, nil
}

func newFirebaseClientForStorage(ctx context.Context) (*firebase.App, error) {
	config := &firebase.Config{StorageBucket: bucketName()}
	app, err := firebase.NewApp(ctx, config)
	return app, err
}
func bucketName() string {
	bucketName := os.Getenv("PROJECT_ID") + ".appspot.com"
	return bucketName
}

func (s *CloudStorage) BucketName() string {
	return bucketName()
}

func (s *CloudStorage) SignedUrl(object string) (string, error) {
	bucket, err := s.client.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("GetDownloadURLUsecase.run: client.DefaultBucket: %v", err)
	}

	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	url, err := bucket.SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("GetDownloadURLUsecase.run: bucket.SignedURL: %v", err)
	}

	return url, nil
}
