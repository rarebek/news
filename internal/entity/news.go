package entity

import "time"

type News struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ImageURL       string    `json:"image_url"`
	CreatedAt      time.Time `json:"created_at"`
	SubCategoryIDs []string  `json:"subcategory_ids"`
}

type NewsWithCategoryNames struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ImageURL        string    `json:"image_url"`
	CreatedAt       time.Time `json:"created_at"`
	CategoryName    string    `json:"category_name"`
	SubCategoryName string    `json:"subcategory_name"`
}

type GetAllNewsRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetNewsBySubCategory struct {
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
	SubCategoryId string `json:"subcategory_id"`
}
