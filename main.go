package main

import (
	"github.com/MCPutro/toko-buku-go/config"
	"github.com/MCPutro/toko-buku-go/controller"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/middleware"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := config.GetConnection()
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	jwtService := service.NewJwtServiceImpl()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, jwtService)
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db)
	bookController := controller.NewBookController(bookService)

	jwtGate := middleware.NewJwtGate(jwtService)

	r := mux.NewRouter()

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/SignUp", userController.SignUp).Methods("POST")
	user.HandleFunc("/SignIn", userController.SignIn).Methods("POST")
	user.HandleFunc("/Books", bookController.FindAll).Methods("GET")

	book := r.PathPrefix("/book").Subrouter()
	book.Use(jwtGate.JwtVerify)
	book.HandleFunc("/Add", bookController.AddBook).Methods("POST")
	book.HandleFunc("/AddStock/{BookId}/{AddStock}", bookController.AddStock).Methods("GET")
	book.HandleFunc("/All", bookController.FindAll).Methods("GET")

	err2 := http.ListenAndServe(":8080", r)
	if err2 != nil {
		helper.PanicIfError(err2)
	}
}
