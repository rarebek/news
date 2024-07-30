package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tarkib.uz/internal/entity"
	"tarkib.uz/pkg/postgres"
)

type AdRepo struct {
	*postgres.Postgres
}

func NewAdRepo(pg *postgres.Postgres) *AdRepo {
	return &AdRepo{pg}
}

func (a *AdRepo) CreateAd(ctx context.Context, request *entity.Ad) error {
	var (
		adID = uuid.NewString()
	)

	data := map[string]interface{}{
		"id":          adID,
		"title":       request.Title,
		"description": request.Description,
		"image_url":   request.ImageURL,
	}
	sql, args, err := a.Builder.Insert("ads").
		SetMap(data).ToSql()
	if err != nil {
		return err
	}

	if _, err = a.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (a *AdRepo) DeleteAd(ctx context.Context, id string) error {
	deleteAdSQL, args, err := a.Builder.Delete("ads").
		Where(squirrel.Eq{
			"id": id,
		}).ToSql()
	if err != nil {
		return err
	}
	if _, err = a.Pool.Exec(ctx, deleteAdSQL, args...); err != nil {
		return err
	}

	return nil
}

func (a *AdRepo) UpdateAd(ctx context.Context, request *entity.Ad) error {
	data := map[string]interface{}{
		"title":       request.Title,
		"description": request.Description,
		"image_url":   request.ImageURL,
	}
	sql, args, err := a.Builder.Update("ads").
		SetMap(data).Where(squirrel.Eq{
		"id": request.ID,
	}).ToSql()
	if err != nil {
		return err
	}

	if _, err = a.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (a *AdRepo) GetAd(ctx context.Context, request *entity.GetAdRequest) (*entity.Ad, error) {
	var (
		ad entity.Ad
	)
	if request.IsAdmin {
		query := a.Builder.Select("id, title, description, image_url, view_count")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, err
		}
		row := a.Pool.QueryRow(ctx, sql, args...)

		if err := row.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.ImageURL, &ad.ViewCount); err != nil {
			return nil, err
		}

		return &ad, nil
	} else {
		query := a.Builder.Select("id, title, description, image_url")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, err
		}
		row := a.Pool.QueryRow(ctx, sql, args...)

		if err := row.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.ImageURL); err != nil {
			return nil, err
		}
		return &ad, nil
	}

}
