package model

import "time"

type TransactionProductModel struct {
	ID int64 `json:"id"`
	TransactionID int64 `json:"transaction_id"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	Price float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
