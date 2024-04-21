package services

import (
	"context"
	"fmt"
	"gateway/api/uploader"
	"gateway/internal/config"
	"gateway/internal/domain"
	"gateway/pkg/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"math"
)

type FileService struct {
	cfg    *config.Config
	client uploader.UploaderClient
}

func NewFileService(cfg *config.Config) domain.FileService {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Services.FileService.Host, cfg.Services.FileService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := uploader.NewUploaderClient(conn)

	return &FileService{
		cfg:    cfg,
		client: client,
	}
}

func (f *FileService) Upload(ctx context.Context, reader file.Reader) (string, error) {
	chunkCount := int(math.Ceil(float64(reader.Size()) / (1 << 20))) // 1MB
	stream, err := f.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	for i := 0; i < chunkCount; i++ {
		chunk := make([]byte, 1<<20)
		_, err := reader.Read(chunk)
		if err != nil {
			return "", err
		}

		err = stream.Send(&uploader.Image{
			Metadata: &uploader.Metadata{
				ContentType: reader.ContentType(),
			},
			Chunk: chunk,
		})
		if err != nil {
			return "", err
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.InvalidArgument {
				return "", fmt.Errorf("invalid image")
			}
		}

		return "", err
	}

	return reply.Id, nil
}
