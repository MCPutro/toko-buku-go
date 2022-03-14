package entity

import "time"

type Transaction struct {
	Id       string `gorm:"primarykey"`
	Date     time.Time
	Customer string
	BookID   string
	Quantity uint8
	Discount uint8
	Total    float32
}
