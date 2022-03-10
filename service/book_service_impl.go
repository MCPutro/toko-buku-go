package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	Repository repository.BookRepository
	DB         *gorm.DB
}

func NewBookService(bookRepo repository.BookRepository, DB *gorm.DB) BookService {
	return &BookServiceImpl{Repository: bookRepo, DB: DB}
}

func (b *BookServiceImpl) AddBook(ctx context.Context, book helper.BookRequest) (*helper.BookResponse, error) {
	newBook := entity.Book{
		Title:    book.Title,
		Author:   book.Author,
		Stock:    book.Stock,
		Price:    book.Price,
		Discount: book.Discount,
	}
	bookId, err := b.Repository.Save(ctx, b.DB, newBook)
	if err != nil {
		return nil, err
	}

	return &helper.BookResponse{
		ID:       bookId,
		Title:    newBook.Title,
		Author:   newBook.Author,
		Stock:    newBook.Stock,
		Price:    newBook.Price,
		Discount: newBook.Discount,
	}, nil
}

func (b *BookServiceImpl) AddStock(ctx context.Context, bookId uint8, newStock uint8) (*helper.BookResponse, error) {
	existingBook, err := b.Repository.FindById(ctx, b.DB, bookId)
	if err != nil {
		return nil, err
	}
	if existingBook == nil {
		return nil, errors.New("book id doesn't exist")
	}

	// if exists than update stock
	existingBook.Stock += newStock
	result, err := b.Repository.Save(ctx, b.DB, *existingBook)
	fmt.Println(result)
	return nil, err
}
func (b *BookServiceImpl) GetListBook(ctx context.Context) (*[]helper.BookResponse, error) {
	books, err := b.Repository.FindAll(ctx, b.DB)
	if err != nil {
		return nil, err
	}

	var tmp []helper.BookResponse

	for _, book := range *books {
		tmp = append(tmp, helper.BookResponse{
			ID:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Stock:    book.Stock,
			Price:    book.Price,
			Discount: book.Discount,
		})
	}

	return &tmp, nil
}
