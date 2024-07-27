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

// func (uc *AdUseCase) GetAdById(ctx context.Context, id string) (*entity.Ad, error) {
// 	if id == "" {
// 		return nil, errors.New("ad ID is required")
// 	}

// 	ad, err := uc.repo.GetAdById(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ad, nil
// }

// func (uc *AdUseCase) GetAllAds(ctx context.Context) ([]entity.Ad, error) {
// 	ads, err := uc.repo.GetAllAds(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ads, nil
// }

func (uc *AdUseCase) DeleteExpiredAds(ctx context.Context) error {
	return uc.repo.DeleteExpiredAds(ctx)
}
