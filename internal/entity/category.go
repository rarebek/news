package entity

type SubCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	SubCategories []SubCategory `json:"subcategories"`
}
