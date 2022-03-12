package entity

type Book struct {
	ID          string `gorm:"primarykey"`
	Title       string
	Author      string
	Stock       uint8
	Price       float32
	Discount    uint8
	Transaction Transaction
}
