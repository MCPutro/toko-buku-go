package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/helper"
)

type BookService interface {
	AddBook(ctx context.Context, book helper.BookRequest) (*helper.BookResponse, error)
	UpdateBook(ctx context.Context, uBook helper.BookRequest, BookId uint8) (*helper.BookResponse, error)
	GetListBook(ctx context.Context) (*[]helper.BookResponse, error)
	DeleteBook(ctx context.Context, bookId uint8) error
	GetBookById(ctx context.Context, bookId uint8) (*helper.BookResponse, error)
}
