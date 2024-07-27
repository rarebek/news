package models

import "time"

type NewsWithCategoryNames struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ImageURL        string    `json:"image_url"`
	CreatedAt       time.Time `json:"created_at"`
	CategoryName    string    `json:"category_name"`
	SubCategoryName string    `json:"subcategory_name"`
}

type News struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ImageURL       string   `json:"image_url"`
	SubCategoryIDs []string `json:"subcategory_ids"`
	Links          []Link   `json:"link"`
}

type Link struct {
	LinkName string `json:"link_name"`
	LinkURL  string `json:"link_url"`
}
