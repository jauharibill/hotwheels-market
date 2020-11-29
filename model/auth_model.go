package model

import "time"

type AuthModel struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
