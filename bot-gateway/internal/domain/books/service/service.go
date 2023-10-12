package books_service

import (
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service struct {
	client pb.BooksServiceClient
}

func NewService(client pb.BooksServiceClient) *Service {
	return &Service{
		client: client,
	}
}
