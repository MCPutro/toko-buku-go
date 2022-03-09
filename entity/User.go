package entity

type UserType int8

const (
	Admin UserType = iota
	Customer
)

type User struct {
	ID          uint8  `gorm:"primarykey"`
	Email       string `gorm:"uniqueIndex"`
	UserName    string
	Password    string
	UserType    UserType
	Transaction Transaction `gorm:"foreignKey:Customer;references:Email"`
}
