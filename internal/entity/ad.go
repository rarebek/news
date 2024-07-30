package entity

type Ad struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ViewCount   int    `json:"view_count"`
}

type CreateAdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type GetAdRequest struct {
	IsAdmin bool
}
