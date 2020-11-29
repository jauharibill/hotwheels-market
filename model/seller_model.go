package model

import "time"

type SellerModel struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
