package controller

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	service service.BookService
}

func NewBookController(service service.BookService) BookController {
	return &BookControllerImpl{service: service}
}

func (b *BookControllerImpl) AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook helper.BookRequest
	helper.ReadFromRequestBody(r, &newBook)

	addBookResponse, err := b.service.AddBook(r.Context(), newBook)

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
			Data:   addBookResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (b *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	//helper.ReadFromRequestBody(r, &newBook)

	book, err := b.service.GetListBook(r.Context())

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
			Data:   book,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (b *BookControllerImpl) AddStock(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookId, _ := strconv.ParseUint(param["BookId"], 10, 8)
	addStock, _ := strconv.ParseUint(param["AddStock"], 10, 8)

	bookResponse, err := b.service.AddStock(r.Context(), uint8(bookId), uint8(addStock))

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
			Data:   bookResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}
