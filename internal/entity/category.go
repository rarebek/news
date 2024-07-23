package entity

type SubCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	SubCategories []SubCategory `json:"subcategories"`
}
