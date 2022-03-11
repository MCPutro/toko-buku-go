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
	"html/template"
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

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository, bookRepository, db)
	transactionController := controller.NewTransactionController(transactionService)

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/SignUp", userController.SignUp).Methods("POST")
	user.HandleFunc("/SignIn", me).Methods("GET")
	user.HandleFunc("/SignIn", userController.SignIn).Methods("POST")
	user.HandleFunc("/Books", bookController.GetListBook).Methods("GET")

	book := r.PathPrefix("/book").Subrouter()
	book.HandleFunc("/Add", bookController.AddBook).Methods("POST")
	book.HandleFunc("/AddStock/{BookId}/{AddStock}", bookController.AddStock).Methods("GET")
	book.HandleFunc("/All", bookController.GetListBook).Methods("GET")
	book.HandleFunc("/Delete/{BookId}", bookController.DeleteBook).Methods("GET")

	r.HandleFunc("/transaction", transactionController.BuyBook).Methods("POST")

	err2 := http.ListenAndServe(":8080", middleware.NewMiddleware(r, jwtService))
	if err2 != nil {
		helper.PanicIfError(err2)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	err := MyTemplates.ExecuteTemplate(w, "home.gohtml", "/user/SignIn")
	if err != nil {
		return
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.gohtml")
	t.Execute(w, nil)
}
