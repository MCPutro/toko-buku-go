package controller

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"net/http"
	"strconv"
)

type TransactionControllerImpl struct {
	service service.TransactionService
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

	transaction, err := t.service.BuyBook(r.Context(), trxRequest)

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
			Data:   transaction,
		}
	}

	if contentType == "application/x-www-form-urlencoded" {
		http.Redirect(w, r, "/listBook", http.StatusSeeOther)
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}

}
