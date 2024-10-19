package repository

import (
	"context"
	"time"
)

type IStorageRepository interface {
	GenerateSignedURL(ctx context.Context, bucketName, objectName string, expires time.Time) (string, error)
}
