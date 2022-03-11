package controller

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"net/http"
)

type TransactionControllerImpl struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &TransactionControllerImpl{service: service}
}

func (t *TransactionControllerImpl) BuyBook(w http.ResponseWriter, r *http.Request) {
	var trxRequest helper.TransactionRequest
	helper.ReadFromRequestBody(r, &trxRequest)

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

	helper.WriteToResponseBody(w, webResponse)

}
