package usecase

import (
	"context"

	"github.com/go-redis/redis/v8"
	"tarkib.uz/config"
	"tarkib.uz/internal/entity"
)

type NewsUseCase struct {
	repo        NewsRepo
	cfg         *config.Config
	RedisClient *redis.Client
}

func NewNewsUseCase(r NewsRepo, cfg *config.Config) *NewsUseCase {
	return &NewsUseCase{
		repo: r,
		cfg:  cfg,
	}
}

func (n *NewsUseCase) CreateNews(ctx context.Context, request *entity.News) error {
	return n.repo.CreateNews(ctx, request)
}

func (n *NewsUseCase) GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error) {
	news, err := n.repo.GetAllNews(ctx, request)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n *NewsUseCase) GetAllNewsByCategory(ctx context.Context, request *entity.GetNewsBySubCategory) ([]entity.NewsWithCategoryNames, error) {
	news, err := n.repo.GetAllNewsByCategory(ctx, request)
	if err != nil {
		return nil, err
	}

	return news, nil
}
