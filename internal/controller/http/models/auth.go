package models

type AdminLoginRequest struct {
	Username string `json:"username" example:"test"`
	Password string `json:"password" example:"test"`
	Avatar   string `json:"avatar"`
}

type Admin struct {
	ID       string `json:"id"`
	Username string `json:"username" example:"test"`
	Password string `json:"password" example:"test"`
	Avatar   string `json:"avatar"`
}

type SuperAdminLoginRequest struct {
	PhoneNumber string `json:"phone_number" example:"test"`
	Password    string `json:"password" example:"test"`
	Avatar      string `json:"avatar"`
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
