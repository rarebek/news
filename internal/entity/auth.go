package entity

type Admin struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
