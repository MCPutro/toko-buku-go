package controller

import "net/http"

type BookController interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	AddStock(w http.ResponseWriter, r *http.Request)
}
