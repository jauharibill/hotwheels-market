package model

type TransactionModel struct {
	ID int64 `json:"id"`
	InvoiceNumber string `json:"invoice"`
	UserID int64 `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
}
