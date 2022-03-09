package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/helper"
)

type BookService interface {
	AddBook(ctx context.Context, book helper.BookRequest) (*helper.BookResponse, error)
	UpdateStock(ctx context.Context, bookId uint8, newStock uint8) (*helper.BookResponse, error)
}
