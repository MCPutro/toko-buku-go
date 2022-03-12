package entity

import "time"

type Transaction struct {
	ID       uint8 `gorm:"primarykey"`
	Date     time.Time
	Customer string
	BookID   uint8
	Quantity uint8
	Discount uint8
	Total    float32
}
