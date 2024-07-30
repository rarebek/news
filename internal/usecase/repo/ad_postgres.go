package repo

import (
	"context"
	"fmt"

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
	var ad entity.Ad

	if request.IsAdmin {
		query := a.Builder.Select("id, title, description, image_url, view_count").From("ads")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, fmt.Errorf("failed to build SQL query: %w", err)
		}

		row := a.Pool.QueryRow(ctx, sql, args...)
		if err := row.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.ImageURL, &ad.ViewCount); err != nil {
			return nil, fmt.Errorf("failed to scan ad for admin: %w", err)
		}

		return &ad, nil
	} else {
		// Update view count for non-admin users
		updateQuery := "UPDATE ads SET view_count = view_count + 1 RETURNING view_count"
		var newViewCount int
		if err := a.Pool.QueryRow(ctx, updateQuery).Scan(&newViewCount); err != nil {
			return nil, fmt.Errorf("failed to update view count: %w", err)
		}

		// Debug log to verify new view count
		fmt.Printf("New view count after update: %d\n", newViewCount)

		query := a.Builder.Select("id, title, description, image_url, view_count").From("ads")
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, fmt.Errorf("failed to build SQL query: %w", err)
		}

		row := a.Pool.QueryRow(ctx, sql, args...)
		if err := row.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.ImageURL, &ad.ViewCount); err != nil {
			return nil, fmt.Errorf("failed to scan ad for non-admin: %w", err)
		}

		return &ad, nil
	}
}
