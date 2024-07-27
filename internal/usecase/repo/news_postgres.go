package repo

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
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

	// Marshal Links to JSON
	linksJSON, err := json.Marshal(request.Links)
	if err != nil {
		return err
	}

	pp.Println(linksJSON)

	data := map[string]interface{}{
		"id":          newsID,
		"name":        request.Name,
		"description": request.Description,
		"image_url":   request.ImageURL,
		"links":       linksJSON, // Store JSONB data
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
	deleteSubcategoryNewsSQL, args, err := n.Builder.Delete("subcategory_news").
		Where(squirrel.Eq{
			"news_id": id,
		}).ToSql()
	if err != nil {
		return err
	}
	if _, err = n.Pool.Exec(ctx, deleteSubcategoryNewsSQL, args...); err != nil {
		return err
	}

	deleteNewsSQL, args, err := n.Builder.Delete("news").
		Where(squirrel.Eq{
			"id": id,
		}).ToSql()
	if err != nil {
		return err
	}
	if _, err = n.Pool.Exec(ctx, deleteNewsSQL, args...); err != nil {
		return err
	}

	return nil
}

func (n *NewsRepo) GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error) {
	var (
		newsList []entity.News
		ids      []string
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
		var linksJSON []byte

		if err := rows.Scan(&news.ID, &news.Name, &news.Description, &news.ImageURL, &news.CreatedAt, &linksJSON); err != nil {
			return nil, err
		}

		pp.Println(string(linksJSON))

		var links []entity.Link
		if err := json.Unmarshal(linksJSON, &links); err != nil {
			return nil, err
		}
		news.Links = links

		sql, args, err = n.Builder.Select("subcategory_id").
			From("subcategory_news").
			Where(squirrel.Eq{
				"news_id": news.ID,
			}).ToSql()
		if err != nil {
			return nil, err
		}

		subCategoryRows, err := n.Pool.Query(ctx, sql, args...)
		if err != nil {
			return nil, err
		}

		for subCategoryRows.Next() {
			var id string
			if err = subCategoryRows.Scan(&id); err != nil {
				return nil, err
			}

			ids = append(ids, id)
		}
		subCategoryRows.Close()

		news.SubCategoryIDs = ids

		newsList = append(newsList, news)

		ids = nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

func (n *NewsRepo) GetFilteredNews(ctx context.Context, request *entity.GetFilteredNewsRequest) ([]entity.News, error) {
	var (
		newsList []entity.News
	)

	queryBuilder := n.Builder.Select("DISTINCT n.id, n.name, n.description, n.image_url, n.created_at, n.links").
		From("news n")

	if len(request.SubCategoryIDs) > 0 {
		queryBuilder = queryBuilder.
			Join("subcategory_news sn ON n.id = sn.news_id").
			Join("subcategories sc ON sn.subcategory_id = sc.id").
			Where(squirrel.Eq{"sc.id": request.SubCategoryIDs})
	}

	if request.CategoryID != "" {
		queryBuilder = queryBuilder.
			Join("subcategory_news sn2 ON n.id = sn2.news_id").
			Join("subcategories sc2 ON sn2.subcategory_id = sc2.id").
			Where(squirrel.Eq{"sc2.category_id": request.CategoryID})
	}

	offset := (request.Page - 1) * request.Limit
	queryBuilder = queryBuilder.
		OrderBy("n.created_at DESC").
		Limit(uint64(request.Limit)).
		Offset(uint64(offset))

	sql, args, err := queryBuilder.ToSql()
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
		var linksJSON []byte

		if err := rows.Scan(&news.ID, &news.Name, &news.Description, &news.ImageURL, &news.CreatedAt, &linksJSON); err != nil {
			return nil, err
		}

		var links []entity.Link
		if err := json.Unmarshal(linksJSON, &links); err != nil {
			return nil, err
		}
		news.Links = links

		subCategoryIDsSQL, subCategoryIDsArgs, err := n.Builder.Select("subcategory_id").
			From("subcategory_news").
			Where(squirrel.Eq{"news_id": news.ID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		subCategoryRows, err := n.Pool.Query(ctx, subCategoryIDsSQL, subCategoryIDsArgs...)
		if err != nil {
			return nil, err
		}

		for subCategoryRows.Next() {
			var id string
			if err = subCategoryRows.Scan(&id); err != nil {
				return nil, err
			}
			news.SubCategoryIDs = append(news.SubCategoryIDs, id)
		}
		subCategoryRows.Close()

		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}
