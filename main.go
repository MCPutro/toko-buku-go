package main

import (
	"context"
	"github.com/MCPutro/toko-buku-go/config"
	"github.com/MCPutro/toko-buku-go/controller"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/middleware"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/MCPutro/toko-buku-go/service"
	t "github.com/MCPutro/toko-buku-go/template"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	db = config.GetConnection()

	jwtService = service.NewJwtServiceImpl()

	userRepository = repository.NewUserRepository()
	userService    = service.NewUserService(userRepository, db, jwtService)
	userController = controller.NewUserController(userService)

	bookRepository = repository.NewBookRepository()
	bookService    = service.NewBookService(bookRepository, db)
	bookController = controller.NewBookController(bookService)

	transactionRepository = repository.NewTransactionRepository()
	transactionService    = service.NewTransactionService(transactionRepository, bookRepository, db)
	transactionController = controller.NewTransactionController(transactionService)
)

func main() {

	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	r := mux.NewRouter()

	restUser := r.PathPrefix("/user").Subrouter()
	restUser.HandleFunc("/SignUp", userController.SignUp).Methods("POST")
	restUser.HandleFunc("/SignIn", userController.SignIn).Methods("POST")
	restUser.HandleFunc("/Books", bookController.GetListBook).Methods("GET")

	restBook := r.PathPrefix("/book").Subrouter()
	restBook.HandleFunc("/Add", bookController.AddBook).Methods("POST")
	restBook.HandleFunc("/AddStock/{BookId}/{AddStock}", bookController.AddStock).Methods("GET")
	restBook.HandleFunc("/All", bookController.GetListBook).Methods("GET")
	restBook.HandleFunc("/Delete/{BookId}", bookController.DeleteBook).Methods("GET")

	r.HandleFunc("/transaction", transactionController.BuyBook).Methods("POST")

	//form ui website
	r.HandleFunc("/", home).Methods(http.MethodGet)
	r.HandleFunc("/login", SignInForm).Methods(http.MethodGet)
	r.HandleFunc("/listBookAdmin", ListBookAdmin).Methods(http.MethodGet)
	r.HandleFunc("/DeleteBookAdmin/{BookId}", DeleteBookAdmin).Methods(http.MethodGet)
	r.HandleFunc("/BookInfoAdmin", BookInfoForm).Methods(http.MethodGet)

	err2 := http.ListenAndServe(":8080", middleware.NewMiddleware(r, jwtService))
	if err2 != nil {
		helper.PanicIfError(err2)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "home.gohtml", "/login")
	if err != nil {
		return
	}
}

func SignInForm(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		return
	}
}

func ListBookAdmin(w http.ResponseWriter, r *http.Request) {
	listBook, _ := bookService.GetListBook(context.Background())

	err := t.MyTemplates.ExecuteTemplate(w, "listBook-admin.gohtml", listBook)
	if err != nil {
		return
	}
}

func DeleteBookAdmin(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookId, _ := strconv.ParseUint(param["BookId"], 10, 8)
	bookService.DeleteBook(r.Context(), uint8(bookId))

	http.Redirect(w, r, "http://"+r.Host+"/listBookAdmin", http.StatusPermanentRedirect)
	//ListBookAdmin(w, r)
}

func BookInfoForm(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "addBook.gohtml", nil)
	if err != nil {
		return
	}
}
