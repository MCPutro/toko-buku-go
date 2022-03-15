package helper

import "time"

type TransactionRequest struct {
	Customer string
	BookID   string
	Quantity uint8
}

type TransactionResponse struct {
	Id        string `gorm:"primarykey"`
	Date      time.Time
	Customer  string
	BookID    string
	BookTitle string
	Price     float32
	Quantity  uint8
	Discount  uint8
	Total     float32
}
