package entity

import "time"

type Transaction struct {
	ID       uint8 `gorm:"primarykey"`
	Date     time.Time
	Customer string
	BookID   uint8
	Quantity float32
	Discount float32
	Total    float32
}
