package books_repository

import (
	"context"
	books_dto "gateway/internal/domain/books/dto"
	"github.com/go-redis/redis"
)

type Repository struct {
	redis *redis.Client
}

func (r *Repository) Create(ctx context.Context, dto books_dto.CreateSearchDTO) error {
	err := r.redis.Set(dto.ID, dto.Searched, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
