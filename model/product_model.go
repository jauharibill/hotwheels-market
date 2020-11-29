package model

import "time"

type ProductModel struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
