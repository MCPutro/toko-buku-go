package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/helper"
)

type BookService interface {
	AddBook(ctx context.Context, book helper.BookRequest) (*helper.BookResponse, error)
	UpdateBook(ctx context.Context, uBook helper.BookRequest, BookId string) (*helper.BookResponse, error)
	UpdateStock(ctx context.Context, BookId string, newStock uint8) error
	GetListBook(ctx context.Context) (*[]helper.BookResponse, error)
	DeleteBook(ctx context.Context, bookId string) error
	GetBookById(ctx context.Context, bookId string) (*helper.BookResponse, error)
}
