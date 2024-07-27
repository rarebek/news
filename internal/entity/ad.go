package entity

import "time"

type Ad struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	ImageURL       string    `json:"image_url"`
	Options        []string  `json:"options"`
	Duration       string    `json:"duration"`
	ExpirationTime time.Time `json:"expiration_time"`
}
