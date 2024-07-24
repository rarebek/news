package entity

type Admin struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type SuperAdmin struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
}

type AdminLoginResponse struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	AccessToken string `json:"access_token"`
}

type SuperAdminLoginResponse struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	AccessToken string `json:"access_token"`
}
