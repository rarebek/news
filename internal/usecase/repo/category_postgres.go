package repo

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"tarkib.uz/internal/entity"
	"tarkib.uz/pkg/postgres"
)

type CategoryRepo struct {
	*postgres.Postgres
}

func NewCategoryRepo(pg *postgres.Postgres) *CategoryRepo {
	return &CategoryRepo{pg}
}

func (n *CategoryRepo) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	query := `
		SELECT c.id, c.name, c.created_at, sc.id, sc.name
		FROM categories c
		LEFT JOIN subcategories sc ON c.id = sc.category_id
		ORDER BY c.created_at, c.id, sc.id
	`

	rows, err := n.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoriesMap := make(map[string]*entity.Category)
	for rows.Next() {
		var categoryID string
		var categoryName string
		var createdAt time.Time
		var subCategoryID sql.NullString
		var subCategoryName sql.NullString

		if err := rows.Scan(&categoryID, &categoryName, &createdAt, &subCategoryID, &subCategoryName); err != nil {
			return nil, err
		}

		if _, exists := categoriesMap[categoryID]; !exists {
			categoriesMap[categoryID] = &entity.Category{
				ID:            categoryID,
				Name:          categoryName,
				CreatedAt:     createdAt,
				SubCategories: []entity.SubCategory{},
			}
		}

		if subCategoryID.Valid && subCategoryName.Valid {
			categoriesMap[categoryID].SubCategories = append(categoriesMap[categoryID].SubCategories, entity.SubCategory{
				ID:   subCategoryID.String,
				Name: subCategoryName.String,
			})
		}
	}

	// Convert the map to a slice
	var categories []entity.Category
	for _, category := range categoriesMap {
		// Sort subcategories for each category
		sort.Slice(category.SubCategories, func(i, j int) bool {
			return category.SubCategories[i].ID < category.SubCategories[j].ID
		})
		categories = append(categories, *category)
	}

	// Sort categories by creation time and then by ID
	sort.Slice(categories, func(i, j int) bool {
		if categories[i].CreatedAt.Equal(categories[j].CreatedAt) {
			return categories[i].ID < categories[j].ID
		}
		return categories[i].CreatedAt.Before(categories[j].CreatedAt)
	})

	return categories, nil
}
