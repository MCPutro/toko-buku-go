package controller

import "net/http"

type TransactionController interface {
	BuyBook(w http.ResponseWriter, r *http.Request)
	GetTransactionListByEmail(w http.ResponseWriter, r *http.Request)
}
