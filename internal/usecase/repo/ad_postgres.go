package repo

import (
	"context"
	"encoding/json"
	"time"

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

	optionsJSON, err := json.Marshal(request.Options)
	if err != nil {
		return err
	}

	duration := getDuration(request.Duration)
	expirationTime := time.Now().Add(duration)

	data := map[string]interface{}{
		"id":              adID,
		"title":           request.Title,
		"description":     request.Description,
		"image_url":       request.ImageURL,
		"options":         optionsJSON,
		"expiration_time": expirationTime,
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

func getDuration(option string) time.Duration {
	switch option {
	case "1 day":
		return 24 * time.Hour
	case "2 days":
		return 48 * time.Hour
	case "3 days":
		return 72 * time.Hour
	case "1 week":
		return 7 * 24 * time.Hour
	case "2 weeks":
		return 14 * 24 * time.Hour
	case "monthly":
		return 30 * 24 * time.Hour
	default:
		return 0
	}
}

func (a *AdRepo) DeleteExpiredAds(ctx context.Context) error {
	deleteExpiredSQL, args, err := a.Builder.Delete("ads").
		Where(squirrel.Lt{
			"expiration_time": time.Now(),
		}).ToSql()
	if err != nil {
		return err
	}
	if _, err = a.Pool.Exec(ctx, deleteExpiredSQL, args...); err != nil {
		return err
	}

	return nil
}
