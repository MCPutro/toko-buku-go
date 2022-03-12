package service

import (
	"context"
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
	newBook := &entity.Book{
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

func (b *BookServiceImpl) UpdateBook(ctx context.Context, uBook helper.BookRequest, BookId uint8) (*helper.BookResponse, error) {
	Book := &entity.Book{
		ID:       BookId,
		Title:    uBook.Title,
		Author:   uBook.Author,
		Stock:    uBook.Stock,
		Price:    uBook.Price,
		Discount: uBook.Discount,
	}

	save, err := b.Repository.Save(ctx, b.DB, Book)

	if err != nil {
		return nil, err
	}

	return &helper.BookResponse{
		ID:       save,
		Title:    uBook.Title,
		Author:   uBook.Author,
		Stock:    uBook.Stock,
		Price:    uBook.Price,
		Discount: uBook.Discount,
	}, nil
}

func (b *BookServiceImpl) GetListBook(ctx context.Context) (*[]helper.BookResponse, error) {
	books, err := b.Repository.FindAll(ctx, b.DB)
	if err != nil {
		return nil, err
	}

	if books == nil {
		return nil, nil
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

func (b *BookServiceImpl) DeleteBook(ctx context.Context, bookId uint8) error {
	err := b.Repository.Delete(ctx, b.DB, bookId)

	if err != nil {
		return err
	}

	return nil
}

func (b *BookServiceImpl) GetBookById(ctx context.Context, bookId uint8) (*helper.BookResponse, error) {
	book, err := b.Repository.FindById(ctx, b.DB, bookId)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, nil
	}

	return &helper.BookResponse{
		ID:       book.ID,
		Title:    book.Title,
		Author:   book.Author,
		Stock:    book.Stock,
		Price:    book.Price,
		Discount: book.Discount,
	}, nil

}
