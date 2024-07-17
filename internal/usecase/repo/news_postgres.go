package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tarkib.uz/internal/entity"
	"tarkib.uz/pkg/postgres"
)

type NewsRepo struct {
	*postgres.Postgres
}

func NewNewsRepo(pg *postgres.Postgres) *NewsRepo {
	return &NewsRepo{pg}
}

func (n *NewsRepo) CreateNews(ctx context.Context, request *entity.News) error {
	var (
		newsID = uuid.NewString()
	)
	data := map[string]interface{}{
		"id":          newsID,
		"name":        request.Name,
		"description": request.Description,
		"image_url":   request.ImageURL,
	}
	sql, args, err := n.Builder.Insert("news").
		SetMap(data).ToSql()
	if err != nil {
		return err
	}

	if _, err = n.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	for _, v := range request.SubCategoryIDs {
		data = map[string]interface{}{
			"subcategory_id": v,
			"news_id":        newsID,
		}

		sql, args, err = n.Builder.Insert("subcategory_news").
			SetMap(data).ToSql()
		if err != nil {
			return err
		}

		if _, err = n.Pool.Exec(ctx, sql, args...); err != nil {
			return err
		}
	}

	return nil
}

func (n *NewsRepo) DeleteNews(ctx context.Context, id string) error {
	var count int
	sql, args, err := n.Builder.Delete("news").
		Where(squirrel.Eq{
			"id": id,
		}).ToSql()
	if err != nil {
		return err
	}

	query := `SELECT COUNT(*) from subcategory_news WHERE news_id = $1`
	if err = n.Pool.QueryRow(ctx, query, id).Scan(&count); err != nil {
		return err
	}

	if _, err = n.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		sql, args, err = n.Builder.Delete("subcategory_news").
			Where(squirrel.Eq{
				"news_id": id,
			}).ToSql()
		if err != nil {
			return err
		}

		if _, err = n.Pool.Exec(ctx, sql, args...); err != nil {
			return err
		}
	}

	return nil
}

func (n *NewsRepo) GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.
	News, error) {
	var (
		newsList []entity.News
		ids      []int
	)
	offset := (request.Page - 1) * request.Limit

	sql, args, err := n.Builder.Select("*").
		From("news").
		OrderBy("created_at DESC").
		Limit(uint64(request.Limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := n.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var news entity.News
		if err := rows.Scan(&news.ID, &news.Name, &news.Description, &news.ImageURL, &news.CreatedAt); err != nil {
			return nil, err
		}

		sql, args, err = n.Builder.Select("subcategory_id").
			From("subcategory_news").
			Where(squirrel.Eq{
				"news_id": news.ID,
			}).ToSql()
		if err != nil {
			return nil, err
		}

		rows, err := n.Pool.Query(ctx, sql, args...)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err != nil {
				return nil, err
			}

			ids = append(ids, id)

		}

		news.SubCategoryIDs = ids

		newsList = append(newsList, news)

		ids = nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}
