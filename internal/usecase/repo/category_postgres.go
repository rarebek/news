package repo

import (
	"context"
	"database/sql"
	"sort"

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

	categoriesMap := make(map[string]*entity.Category)
	for rows.Next() {
		var categoryID string
		var categoryName string
		var subCategoryID sql.NullString
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
				ID:   subCategoryID.String,
				Name: subCategoryName.String,
			})
		}
	}

	// Define the exact order based on the provided list
	order := []string{
		"b1d7357d-1d92-4e77-a8d0-1d394c5b2ef2",
		"99d0724b-2d49-4140-8b60-0c41fe214b82",
		"09e6dfb7-59e4-4c88-a8b1-97c8906f9c9d",
		"f7f3c65c-09f5-4f6a-b43a-7cfdb60cf5a1",
		"9d57d482-6db2-43ec-8513-6478c066aa51",
		"23f2a6b3-9e47-46e7-b4d6-3b08b8d805f3",
		"6f4a4be8-0b9c-4fa7-b09e-5976e0f43cfb",
		"e3b9b8c2-0f6f-4db9-9c3c-5b9d3e767d1e",
		"72a5ec87-44ed-438c-b9d0-0e391ed7da4d",
		"b0a5e2f4-5f41-451a-9a0a-d7b8c8b53a69",
		"a1e67972-b9b0-4eec-a51d-8fa8d08e51bb",
		"7e27d7bb-258d-4df5-810d-9b5c3146a606",
	}

	// Convert the map to a slice and sort it based on the defined order
	var categories []entity.Category
	for _, id := range order {
		if category, exists := categoriesMap[id]; exists {
			// Sort subcategories for each category
			sort.Slice(category.SubCategories, func(i, j int) bool {
				return category.SubCategories[i].ID < category.SubCategories[j].ID
			})
			categories = append(categories, *category)
		}
	}

	return categories, nil
}
