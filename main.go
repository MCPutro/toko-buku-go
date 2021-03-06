package main

import (
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

	authCheck = middleware.NewAuthCheck(jwtService)
)

func main() {

	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	r := mux.NewRouter()

	restUser := r.PathPrefix("/user").Subrouter()
	restUser.HandleFunc("/SignUp", userController.SignUp).Methods(http.MethodPost)
	restUser.HandleFunc("/SignIn", userController.SignIn).Methods(http.MethodPost)
	//restUser.HandleFunc("/Books", bookController.GetListBook).Methods(http.MethodGet)

	restBook := r.PathPrefix("/book").Subrouter()
	restBook.Use(authCheck.AuthAdmin)
	restBook.HandleFunc("/Add", bookController.AddBook).Methods(http.MethodPost)
	restBook.HandleFunc("/All", bookController.GetListBook).Methods(http.MethodGet)
	restBook.HandleFunc("/{BookId}", bookController.GetBookById).Methods(http.MethodGet)
	restBook.HandleFunc("/Update/{BookId}", bookController.UpdateBook).Methods(http.MethodPost)
	restBook.HandleFunc("/Delete/{BookId}", bookController.DeleteBook).Methods(http.MethodGet)

	r.Handle("/transaction", authCheck.AuthCustomer(http.HandlerFunc(transactionController.BuyBook))).Methods(http.MethodPost)
	r.Handle("/transaction/history/{Email}", authCheck.AuthCustomer(http.HandlerFunc(transactionController.GetTransactionListByEmail))).Methods(http.MethodGet)

	//form ui website
	r.HandleFunc("/", home).Methods(http.MethodGet)
	r.HandleFunc("/login", SignInForm).Methods(http.MethodGet)
	r.HandleFunc("/SignUp", SignUpForm).Methods(http.MethodGet)
	//admin
	r.Handle("/listBookAdmin", authCheck.AuthAdmin(http.HandlerFunc(ListBookAdmin))).Methods(http.MethodGet)
	r.Handle("/DeleteBookAdmin/{BookId}", authCheck.AuthAdmin(http.HandlerFunc(DeleteBookAdmin))).Methods(http.MethodGet)
	r.Handle("/AddBookFormAdmin", authCheck.AuthAdmin(http.HandlerFunc(AddBookForm))).Methods(http.MethodGet)
	r.Handle("/BookInfoFormAdmin/{BookId}", authCheck.AuthAdmin(http.HandlerFunc(BookInfoFormAdmin))).Methods(http.MethodGet)
	//customer
	r.Handle("/listBook", authCheck.AuthCustomer(http.HandlerFunc(ListBook))).Methods(http.MethodGet)
	r.Handle("/buy/{BookId}", authCheck.AuthCustomer(http.HandlerFunc(BuyBook))).Methods(http.MethodGet)
	r.Handle("/transaction/history", authCheck.AuthCustomer(http.HandlerFunc(transactionController.GetTransactionListByEmail))).Methods(http.MethodGet)

	err2 := http.ListenAndServe(":8080", r) //http.ListenAndServe(":8080", middleware.NewMiddleware(r, jwtService))
	if err2 != nil {
		helper.PanicIfError(err2)
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"signIn": "/login",
		"signUp": "/SignUp",
	}

	err := t.MyTemplates.ExecuteTemplate(w, "home.gohtml", data)
	if err != nil {
		return
	}
}

func SignInForm(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "signIn.gohtml", nil)
	if err != nil {
		return
	}
}

func SignUpForm(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "signUp.gohtml", nil)
	if err != nil {
		return
	}
}

func ListBookAdmin(w http.ResponseWriter, r *http.Request) {
	listBook, _ := bookService.GetListBook(r.Context())

	data := map[string]interface{}{
		"Email": helper.GetCookie(r, "email"),
		"Books": listBook,
	}

	err := t.MyTemplates.ExecuteTemplate(w, "listBook-admin.gohtml", data)
	if err != nil {
		return
	}
}

func DeleteBookAdmin(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	err := bookService.DeleteBook(r.Context(), param["BookId"])
	if err != nil {
		return
	}

	http.Redirect(w, r, "/listBookAdmin", http.StatusPermanentRedirect)
	//ListBookAdmin(w, r)
}

func AddBookForm(w http.ResponseWriter, r *http.Request) {
	err := t.MyTemplates.ExecuteTemplate(w, "addBook.gohtml", nil)
	if err != nil {
		return
	}
}

func BookInfoFormAdmin(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookById, err := bookService.GetBookById(r.Context(), param["BookId"])

	if err != nil {
		return
	}

	err2 := t.MyTemplates.ExecuteTemplate(w, "bookInfo.gohtml", bookById)
	if err2 != nil {
		return
	}
}

func ListBook(w http.ResponseWriter, r *http.Request) {
	listBook, _ := bookService.GetListBook(r.Context())

	data := map[string]interface{}{
		"Email": helper.GetCookie(r, "email"),
		"Books": listBook,
	}

	err := t.MyTemplates.ExecuteTemplate(w, "listBook-customer.gohtml", data)
	if err != nil {
		return
	}
}

func BuyBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookById, err := bookService.GetBookById(r.Context(), param["BookId"])

	if err != nil {
		return
	}

	err2 := t.MyTemplates.ExecuteTemplate(w, "buyBook.gohtml", bookById)
	if err2 != nil {
		return
	}

}
