package models

type AdminLoginRequest struct {
	PhoneNumber string `json:"phone_number" example:"+998889561006"`
	Password    string `json:"password" example:"Nodirbek"`
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
