package grpc

import (
	"bufio"
	"context"
	"errors"
	"file-service/api/uploader"
	"file-service/internal/entities"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
)

const (
	BUCKET_DEFAULT  = "default"
	BUCKET_PANORAMA = "panoramas"
)

type Service interface {
	Upload(ctx context.Context, file *entities.File, bucket string) error
}

type Handler struct {
	service Service
	uploader.UnimplementedUploaderServer
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Upload(stream uploader.Uploader_UploadServer) error {
	file := &entities.File{}
	writer := bufio.NewWriter(&file.Buffer)
	log := slog.With("handler", "upload")

	for {
		object, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Error("error occurred while uploading", slog.String("err", err.Error()))
			return status.Error(codes.Internal, fmt.Sprintf("%s", err.Error()))
		}

		chunk := object.Chunk

		if file.ContentType == "" {
			log.Debug("setting content type", slog.String("contentType", object.Metadata.ContentType))
			file.ContentType = object.Metadata.ContentType
		}

		file.Size += int64(len(chunk))
		if _, err := writer.Write(chunk); err != nil {
			fmt.Printf("%s", err.Error())
			return err
		}
		log.Debug("received bytes", slog.Int("chunkSize", len(chunk)))
	}

	err := h.service.Upload(stream.Context(), file, BUCKET_DEFAULT)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&uploader.ImageInfo{
		Id: file.Name,
	})
}

func (h *Handler) UploadPanorama(stream uploader.Uploader_UploadPanoramaServer) error {

	file := &entities.File{}
	writer := bufio.NewWriter(&file.Buffer)
	log := slog.With("handler", "uploadPanorama")

	for {
		object, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Error("error occurred while uploading", slog.String("err", err.Error()))
			return status.Error(codes.Internal, fmt.Sprintf("%s", err.Error()))
		}

		chunk := object.Chunk

		if file.ContentType == "" {
			log.Debug("setting content type", slog.String("contentType", object.Metadata.ContentType))
			file.ContentType = object.Metadata.ContentType
		}

		file.Size += int64(len(chunk))
		if _, err := writer.Write(chunk); err != nil {
			fmt.Printf("%s", err.Error())
			return err
		}
		log.Debug("received bytes", slog.Int("chunkSize", len(chunk)))
	}

	err := h.service.Upload(stream.Context(), file, BUCKET_PANORAMA)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&uploader.ImageInfo{
		Id: file.Name,
	})
}
