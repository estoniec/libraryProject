package booksService

import (
	"context"
	"encoding/json"
	books_dto "gateway/internal/domain/books/dto"
	"github.com/go-redis/redis"
	"log/slog"
)

type Storage struct {
	redis *redis.Client
}

func NewBooksStorage(redis *redis.Client) *Storage {
	return &Storage{
		redis: redis,
	}
}

func (r *Storage) Create(ctx context.Context, dto books_dto.CreateSearchDTO) error {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = r.redis.SAdd(dto.ID, jsonData).Err()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}

func (r *Storage) Find(ctx context.Context, dto books_dto.FindSearchDTO) (books_dto.FindSearchOutput, error) {
	var output books_dto.FindSearchOutput
	isExists, err := r.redis.Exists(dto.ID).Result()
	if err != nil {
		slog.Error(err.Error())
		return output, err
	}
	if isExists == 0 {
		return output, nil
	}
	jsonData, err := r.redis.SMembers(dto.ID).Result()
	if err != nil {
		slog.Error(err.Error())
		return output, err
	}
	err = json.Unmarshal([]byte(jsonData[0]), &output)
	if err != nil {
		slog.Error(err.Error())
		return output, err
	}
	err = r.redis.Del(dto.ID).Err()
	if err != nil {
		slog.Error(err.Error())
		return output, err
	}
	return output, nil
}
