package grpc

import (
	"context"
	"students-service/api/students"
)

type Service interface {
	Create(context.Context, *students.CreateStudentRequest) (*students.CreateStudentResponse, error)
	PostReview(context.Context, *students.PostReviewRequest) (*students.PostReviewResponse, error)
}

type Handler struct {
	students.UnimplementedStudentsServer
}

func (h Handler) Create(ctx context.Context, request *students.CreateStudentRequest) (*students.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostReview(ctx context.Context, request *students.CreateReviewRequest) (*students.CreateReviewResponse, error) {
	//TODO implement me
	panic("implement me")
}
