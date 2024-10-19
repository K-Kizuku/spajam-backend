package repository

import (
	"context"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/K-Kizuku/spajam-backend/internal/domain/repository"
	"github.com/K-Kizuku/spajam-backend/pkg/errors"
)

type StorageRepository struct {
	client *storage.Client
}

func NewStorageRepository(client *storage.Client) repository.IStorageRepository {
	return &StorageRepository{
		client: client,
	}
}

func (r *StorageRepository) GenerateSignedURL(ctx context.Context, bucketName, objectName string, expires time.Time) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "PUT",
		Expires: expires,
		Headers: []string{
			"Content-Type:audio/ogg",
		},
		ContentType: "audio/ogg",
	}

	url, err := r.client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err)
	}
	return url, nil
}
