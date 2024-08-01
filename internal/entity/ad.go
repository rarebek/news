package entity

type Ad struct {
	ID        string `json:"id"`
	Link      string `json:"link"`
	ImageURL  string `json:"image_url"`
	ViewCount int    `json:"view_count"`
}

type CreateAdRequest struct {
	Link     string `json:"link"`
	ImageURL string `json:"image_url"`
}

type GetAdRequest struct {
	IsAdmin bool
}
