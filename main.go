package main

import (
	"context"
	"fmt"
	"github.com/MCPutro/toko-buku-go/config"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/MCPutro/toko-buku-go/service"
)

func main() {
	db := config.GetConnection()
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	transactionRepository := repository.NewTransactionRepository()
	bookRepository := repository.NewBookRepository()

	transactionService := service.NewTransactionService(transactionRepository, bookRepository, db)

	nt := helper.TransactionRequest{
		Customer: "email2@gmail.com",
		BookID:   7,
		Quantity: 10,
	}

	transactionService.BuyBook(context.Background(), nt)
}

func main5() {
	db := config.GetConnection()
	//db.AutoMigrate(&User{}, &CreditCard{})
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db)

	_, err = bookService.UpdateStock(context.Background(), 7, 255)
	fmt.Println(">>", err)
}

func main4() {
	db := config.GetConnection()
	//db.AutoMigrate(&User{}, &CreditCard{})
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db)

	newBook := helper.BookRequest{
		Title:    "buku makan",
		Author:   "yuk",
		Stock:    10,
		Price:    100.9,
		Discount: 0.30,
	}

	book, err := bookService.AddBook(context.Background(), newBook)
	fmt.Println(book)
}

func main3() {
	db := config.GetConnection()
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)

	userCreateRequest := helper.UserCreateRequest{
		Email:    "email2@gmail.com",
		UserName: "email2-username",
		Password: "passss",
		UserType: entity.Customer,
	}

	response, err := userService.CreateNewUser(context.Background(), userCreateRequest)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

	//loginRequest := helper.UserLoginRequest{
	//	Email:    "email2@gmail.com",
	//	Password: "passsss",
	//}
	//
	//login, err := userService.Login(context.Background(), loginRequest)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(login)
	//}

}

func main2() {
	db := config.GetConnection()
	//db.AutoMigrate(&User{}, &CreditCard{})
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	//	userRepository := repository.NewUserRepository()
	//id, err := userRepository.Save(context.Background(), db, entity.User{
	//	Email:    "email1@gmail.com",
	//	UserName: "email-test",
	//	Password: "email-test1",
	//	UserType: entity.Customer,
	//})
	//
	//if err != nil {
	//	fmt.Println("??", err)
	//}
	//
	//fmt.Println(id)
	//fmt.Println("-------------------------")
	//user, err := userRepository.FindByEmail(context.Background(), db, "email1@gmail.com")
	//if err != nil {
	//
	//}
	//
	//fmt.Println(user.Transaction)

	bookRepository := repository.NewBookRepository()

	//book := entity.Book{
	//	Title:    "jalan22",
	//	Stock:    10,
	//	Price:    12,
	//	Discount: 0,
	//	Author:   "kicing",
	//}
	//save, err := bookRepository.Save(context.Background(), db, book)
	//
	//fmt.Println(save)

	//byTitleAndAuthor, err := bookRepository.FindByTitleAndAuthor(context.Background(), db, book.Title, book.Author)
	//
	//fmt.Println(byTitleAndAuthor)
	//
	//all, err := bookRepository.FindAll(context.Background(), db)
	//fmt.Println(all)

	findById, err := bookRepository.FindById(context.Background(), db, 49)

	if err != nil {
		fmt.Println("ini error :", err)
	}

	fmt.Println(findById)
}
