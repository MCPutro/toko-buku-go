package controller

import "net/http"

type BookController interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	GetListBook(w http.ResponseWriter, r *http.Request)
	AddStock(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
