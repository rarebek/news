package repo

import (
	"context"
	ssq "database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"tarkib.uz/internal/entity"
	"tarkib.uz/pkg/postgres"
)

type AdRepo struct {
	*postgres.Postgres
}

func NewAdRepo(pg *postgres.Postgres) *AdRepo {
	return &AdRepo{pg}
}

func (a *AdRepo) adExists(ctx context.Context) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM ads"
	err := a.Pool.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *AdRepo) CreateAd(ctx context.Context, request *entity.Ad) error {
	exists, err := a.adExists(ctx)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("an ad already exists")
	}

	var (
		adID = uuid.NewString()
	)

	data := map[string]interface{}{
		"id":        adID,
		"link":      request.Link,
		"image_url": request.ImageURL,
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
		"link":      request.Link,
		"image_url": request.ImageURL,
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
		var viewCount ssq.NullInt64
		query := a.Builder.Select("id, link, image_url, view_count").From("ads").Limit(1)
		sql, args, err := query.ToSql()
		if err != nil {
			return nil, fmt.Errorf("failed to build SQL query: %w", err)
		}

		row := a.Pool.QueryRow(ctx, sql, args...)
		if err := row.Scan(&ad.ID, &ad.Link, &ad.ImageURL, &viewCount); err != nil {
			return nil, fmt.Errorf("failed to scan ad for admin: %w", err)
		}

		if viewCount.Valid {
			ad.ViewCount = int(viewCount.Int64)
		} else {
			ad.ViewCount = 0
		}

		return &ad, nil
	} else {
		selectQuery := a.Builder.Select("id, link, image_url, view_count").From("ads").Limit(1)
		sql, args, err := selectQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("failed to build SQL query: %w", err)
		}

		row := a.Pool.QueryRow(ctx, sql, args...)
		var viewCount ssq.NullInt64
		if err := row.Scan(&ad.ID, &ad.Link, &ad.ImageURL, &viewCount); err != nil {
			return nil, fmt.Errorf("failed to scan ad for non-admin: %w", err)
		}

		pp.Println("SCANNED AD: ", ad)

		if viewCount.Valid {
			ad.ViewCount = int(viewCount.Int64)
		} else {
			ad.ViewCount = 0
		}

		ad.ViewCount += 1

		updateQuery := "UPDATE ads SET view_count = $1"
		_, err = a.Pool.Exec(ctx, updateQuery, ad.ViewCount)
		if err != nil {
			return nil, fmt.Errorf("failed to execute update query: %w", err)
		}

		return &ad, nil
	}
}
