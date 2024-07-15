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
	var newsID = uuid.NewString()
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

	data = map[string]interface{}{
		"subcategory_id": request.SubCategoryID,
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

	return nil
}

func (n *NewsRepo) GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error) {
	var newsList []entity.News
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
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}
func (n *NewsRepo) GetAllNewsByCategory(ctx context.Context, request *entity.GetNewsBySubCategory) ([]entity.NewsWithCategoryNames, error) {
	var newsList []entity.NewsWithCategoryNames
	offset := (request.Page - 1) * request.Limit

	sql, args, err := n.Builder.
		Select("n.id, n.name, n.description, n.image_url, n.created_at, s.name AS subcategory_name, c.name AS category_name").
		From("news n").
		Join("subcategory_news sn ON n.id = sn.news_id").
		Join("subcategories s ON sn.subcategory_id = s.id").
		Join("categories c ON s.category_id = c.id").
		Where(squirrel.Eq{
			"s.id": request.SubCategoryId,
		}).
		OrderBy("n.created_at DESC").
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
		var news entity.NewsWithCategoryNames
		if err := rows.Scan(&news.ID, &news.Name, &news.Description, &news.ImageURL, &news.CreatedAt, &news.SubCategoryName, &news.CategoryName); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}
