package usecase

import (
	"context"

	"github.com/go-redis/redis/v8"
	"tarkib.uz/config"
	"tarkib.uz/internal/entity"
)

type CategoryUseCase struct {
	repo        CategoryRepo
	cfg         *config.Config
	RedisClient *redis.Client
}

func NewCategoryUseCase(r CategoryRepo, cfg *config.Config) *CategoryUseCase {
	return &CategoryUseCase{
		repo: r,
		cfg:  cfg,
	}
}

func (n *CategoryUseCase) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return n.repo.GetAllCategories(ctx)
}
