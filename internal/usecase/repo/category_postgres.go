package repo

import (
	"context"
	"database/sql"

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
		SELECT c.id, c.name, sc.id, sc.name
		FROM categories c
		LEFT JOIN subcategories sc ON c.id = sc.category_id
		ORDER BY c.id, sc.id
	`

	rows, err := n.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoriesMap := make(map[int]*entity.Category)
	for rows.Next() {
		var categoryID int
		var categoryName string
		var subCategoryID sql.NullInt32
		var subCategoryName sql.NullString

		if err := rows.Scan(&categoryID, &categoryName, &subCategoryID, &subCategoryName); err != nil {
			return nil, err
		}

		if _, exists := categoriesMap[categoryID]; !exists {
			categoriesMap[categoryID] = &entity.Category{
				ID:            categoryID,
				Name:          categoryName,
				SubCategories: []entity.SubCategory{},
			}
		}

		if subCategoryID.Valid && subCategoryName.Valid {
			categoriesMap[categoryID].SubCategories = append(categoriesMap[categoryID].SubCategories, entity.SubCategory{
				ID:   int(subCategoryID.Int32),
				Name: subCategoryName.String,
			})
		}
	}

	var categories []entity.Category
	for _, category := range categoriesMap {
		categories = append(categories, *category)
	}

	return categories, nil
}
