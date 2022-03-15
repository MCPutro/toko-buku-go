package controller

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/MCPutro/toko-buku-go/template"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TransactionControllerImpl struct {
	service     service.TransactionService
	bookService service.BookService
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &TransactionControllerImpl{service: service}
}

func (t *TransactionControllerImpl) BuyBook(w http.ResponseWriter, r *http.Request) {
	var trxRequest helper.TransactionRequest
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			helper.WriteToResponseBody(w, err)
			return
		}

		emailCookie := helper.GetCookie(r, "email")

		if emailCookie == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		Qty, _ := strconv.ParseUint(r.PostFormValue("Qty"), 10, 8)

		trxRequest = helper.TransactionRequest{
			Customer: emailCookie,
			BookID:   r.PostFormValue("BookId"),
			Quantity: uint8(Qty),
		}
	} else {
		helper.ReadFromRequestBody(r, &trxRequest)
	}

	transaction, errBuyBook := t.service.BuyBook(r.Context(), trxRequest)

	var webResponse helper.Response

	if errBuyBook != nil {
		webResponse = helper.Response{
			Status:  "error",
			Message: errBuyBook.Error(),
			Data:    nil,
		}
	} else {
		webResponse = helper.Response{
			Status: "success",
			Data:   transaction,
		}
	}

	if contentType == "application/x-www-form-urlencoded" {
		http.Redirect(w, r, "/listBook", http.StatusSeeOther)
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}

}

func (t *TransactionControllerImpl) GetTransactionListByEmail(w http.ResponseWriter, r *http.Request) {
	//contentType := r.Header.Get("Content-Type")
	st := r.Header.Get("Sec-Ch-Ua")

	var email string
	if st != "" {
		email = helper.GetCookie(r, "email")
	} else {
		param := mux.Vars(r)
		email = param["Email"]
	}

	trxByCustomer, err := t.service.FindByCustomerEmail(r.Context(), email)

	var webResponse helper.Response

	if err != nil {
		webResponse = helper.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		webResponse = helper.Response{
			Status: "success",
			Data:   trxByCustomer,
		}
	}

	if st != "" {
		err := template.MyTemplates.ExecuteTemplate(w, "transactions.gohtml", trxByCustomer)
		if err != nil {
			return
		}
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}

}
