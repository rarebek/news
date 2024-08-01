package entity

type Ad struct {
	ID        string `json:"id"`
	Link      string `json:"link"`
	ImageURL  string `json:"image_url"`
	ViewCount int    `json:"view_count"`
}

// comment
type CreateAdRequest struct {
	Link     string `json:"link"`
	ImageURL string `json:"image_url"`
	ID       string `json:"id"`
}

type GetAdRequest struct {
	IsAdmin bool
	ID      string `json:"id"`
}
