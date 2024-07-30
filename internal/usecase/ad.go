package usecase

import (
	"context"
	"errors"

	"tarkib.uz/config"
	"tarkib.uz/internal/entity"
	"tarkib.uz/internal/usecase/repo"
)

type AdUseCase struct {
	repo repo.AdRepo
}

func NewAdUseCase(repo repo.AdRepo, cfg config.Config) *AdUseCase {
	return &AdUseCase{
		repo: repo,
	}
}

func (uc *AdUseCase) CreateAd(ctx context.Context, ad *entity.Ad) error {
	if ad.Title == "" || ad.Description == "" {
		return errors.New("title and description are required")
	}

	return uc.repo.CreateAd(ctx, ad)
}

func (uc *AdUseCase) DeleteAd(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("ad ID is required")
	}

	return uc.repo.DeleteAd(ctx, id)
}

func (uc *AdUseCase) UpdateAd(ctx context.Context, ad *entity.Ad) error {
	if ad.Title == "" || ad.Description == "" {
		return errors.New("title and description are required")
	}

	return uc.repo.UpdateAd(ctx, ad)
}

func (uc *AdUseCase) GetAd(ctx context.Context, request *entity.GetAdRequest) (*entity.Ad, error) {
	return uc.repo.GetAd(ctx, request)
}
