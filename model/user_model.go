package model

import "time"

type UserModel struct {
	ID int64 `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Country string `json:"country"`
	Phone int `json:"phone"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
