package models

type AdminLoginRequest struct {
	Username string `json:"username" example:"test"`
	Password string `json:"password" example:"test"`
}

type SuperAdminLoginRequest struct {
	PhoneNumber string `json:"phone_number" example:"test"`
	Password    string `json:"password" example:"test"`
}

type AdminLoginResponse struct {
	AccessToken string `json:"access_token"`
}

type Message struct {
	Message string `json:"message"`
}

type SubCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	SubCategories []SubCategoryResponse `json:"subcategories"`
}
