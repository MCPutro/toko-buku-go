package controller

import (
	"fmt"
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
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			helper.WriteToResponseBody(w, err)
			return
		}

		s, _ := strconv.ParseUint(r.PostFormValue("Stock"), 10, 8)
		p, _ := strconv.ParseFloat(r.PostFormValue("Price"), 32)
		d, _ := strconv.ParseFloat(r.PostFormValue("Discount"), 32)

		newBook = helper.BookRequest{
			Title:    r.PostFormValue("Title"),
			Author:   r.PostFormValue("Author"),
			Stock:    uint8(s),
			Price:    float32(p),
			Discount: uint8(d),
		}
	} else {
		helper.ReadFromRequestBody(r, &newBook)
	}

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

	if contentType == "application/x-www-form-urlencoded" {
		http.Redirect(w, r, "/listBookAdmin", http.StatusSeeOther)
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}
}

func (b *BookControllerImpl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookId, _ := strconv.ParseUint(param["BookId"], 10, 8)

	var uBook helper.BookRequest

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			helper.WriteToResponseBody(w, err)
			return
		}

		s, _ := strconv.ParseUint(r.PostFormValue("Stock"), 10, 8)
		p, _ := strconv.ParseFloat(r.PostFormValue("Price"), 32)
		d, _ := strconv.ParseFloat(r.PostFormValue("Discount"), 32)

		uBook = helper.BookRequest{

			Title:    r.PostFormValue("Title"),
			Author:   r.PostFormValue("Author"),
			Stock:    uint8(s),
			Price:    float32(p),
			Discount: uint8(d),
		}
	} else {
		helper.ReadFromRequestBody(r, &uBook)
	}

	fmt.Println("uBook : ", uBook)

	bookResponse, err := b.service.UpdateBook(r.Context(), uBook, uint8(bookId))

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

	if contentType == "application/x-www-form-urlencoded" {
		http.Redirect(w, r, "/listBookAdmin", http.StatusSeeOther)
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}

}

func (b *BookControllerImpl) GetListBook(w http.ResponseWriter, r *http.Request) {

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

func (b *BookControllerImpl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookId, _ := strconv.ParseUint(param["BookId"], 10, 8)

	err := b.service.DeleteBook(r.Context(), uint8(bookId))

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
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (b *BookControllerImpl) GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	bookId, _ := strconv.ParseUint(param["BookId"], 10, 8)

	bookById, err := b.service.GetBookById(r.Context(), uint8(bookId))

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
			Data:   bookById,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}
