package storage

import (
	"context"
	"file-service/internal/config"
	"file-service/internal/entities"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log/slog"
)

type MinioStorage struct {
	client *minio.Client
}

func New(client *minio.Client, cfg *config.Config) (*MinioStorage, error) {
	return &MinioStorage{
		client: client,
	}, nil
}

func (s *MinioStorage) checkBucket(ctx context.Context, bucket string) error {

	log := slog.With("method", "checkBucket").With("bucket", bucket)

	exists, err := s.client.BucketExists(ctx, bucket)
	if err != nil {
		log.Error("unable to check bucket exists", slog.String("err", err.Error()))
		return fmt.Errorf("MinioStorage cannot check bucket existance: %w", err)
	}

	if !exists {
		log.Debug("bucket not exists")
		if err := s.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			log.Error("unable to create bucket", slog.String("err", err.Error()))
			return fmt.Errorf("MinioStorage cannot create bucket: %w", err)
		}
		log.Debug("bucket created")
	}

	return nil
}

func (s *MinioStorage) Create(ctx context.Context, fileInfo *entities.FileInfo, reader io.Reader, bucket string) error {

	log := slog.With("method", "create").With("bucket", bucket).With("fileName", fileInfo.Name)

	if err := s.checkBucket(ctx, bucket); err != nil {
		log.Error("unable to check bucket", slog.String("err", err.Error()))
		return fmt.Errorf("MinioStorage.Create: %w", err)
	}

	log.Debug("putting file")
	_, err := s.client.PutObject(ctx, bucket, fileInfo.Name, reader, fileInfo.Size, minio.PutObjectOptions{
		ContentType: fileInfo.ContentType,
	})
	if err != nil {
		log.Error("unable to put object", slog.String("err", err.Error()))
		return fmt.Errorf("MinioStorage.Create unable to put object: %w", err)
	}

	log.Debug("object putted")

	return nil
}
