package service

import (
	"context"
	"file-service/internal/entities"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/google/uuid"
)

type Storage interface {
	Create(ctx context.Context, fileInfo *entities.FileInfo, reader io.Reader, bucket string) error
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Upload(ctx context.Context, file *entities.File, bucket string) error {

	log := slog.With("method", "upload")

	id, err := uuid.NewV7()
	if err != nil {
		log.Error("unable to create uuid", slog.String("err", err.Error()))
		return fmt.Errorf("Service.Create unable to create UUID: %w", err)
	}

	file.Name = fmt.Sprintf("%s.%s", id.String(), strings.Split(file.ContentType, "/")[1])

	log.Debug("uploading a file", slog.String("filename", file.Name))
	if err := s.storage.Create(ctx, &file.FileInfo, &file.Buffer, bucket); err != nil {
		log.Error("unable to create file", slog.String("err", err.Error()))
		return fmt.Errorf("Service.Create unable to create file: %w", err)
	}

	return nil
}
