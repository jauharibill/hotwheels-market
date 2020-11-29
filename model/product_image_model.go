package model

import "time"

type ProductImageModel struct {
	ID int64 `json:"id"`
	Image string `json:"image"`
	Alter string `json:"alter"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
