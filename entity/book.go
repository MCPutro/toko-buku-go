package entity

type Book struct {
	ID          uint8 `gorm:"primarykey"`
	Title       string
	Author      string
	Stock       uint8
	Price       float32
	Discount    float32
	Transaction Transaction
}
