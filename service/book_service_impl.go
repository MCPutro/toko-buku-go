package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/google/uuid"
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
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)

	newBook := &entity.Book{
		ID:       uuid.New().String(),
		Title:    book.Title,
		Author:   book.Author,
		Stock:    book.Stock,
		Price:    book.Price,
		Discount: book.Discount,
	}
	nBook, err := b.Repository.Save(ctx, trx, newBook)
	if err != nil {
		return nil, err
	}

	return &helper.BookResponse{
		Id:       nBook.ID,
		Title:    nBook.Title,
		Author:   nBook.Author,
		Stock:    nBook.Stock,
		Price:    nBook.Price,
		Discount: nBook.Discount,
	}, nil
}

func (b *BookServiceImpl) UpdateBook(ctx context.Context, uBook helper.BookRequest, BookId string) (*helper.BookResponse, error) {
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)

	Book := &entity.Book{
		ID:       BookId,
		Title:    uBook.Title,
		Author:   uBook.Author,
		Stock:    uBook.Stock,
		Price:    uBook.Price,
		Discount: uBook.Discount,
	}

	bookSaved, err := b.Repository.Save(ctx, trx, Book)

	if err != nil {
		return nil, err
	}

	return &helper.BookResponse{
		Id:       bookSaved.ID,
		Title:    bookSaved.Title,
		Author:   bookSaved.Author,
		Stock:    bookSaved.Stock,
		Price:    bookSaved.Price,
		Discount: bookSaved.Discount,
	}, nil
}

func (b *BookServiceImpl) UpdateStock(ctx context.Context, BookId string, newStock uint8) error {
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)
	return b.Repository.UpdateStock(ctx, trx, BookId, newStock)
}

func (b *BookServiceImpl) GetListBook(ctx context.Context) (*[]helper.BookResponse, error) {
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)

	books, err := b.Repository.FindAll(ctx, trx)
	if err != nil {
		return nil, err
	}

	if books == nil {
		return nil, nil
	}

	var tmp []helper.BookResponse

	for _, book := range *books {
		tmp = append(tmp, helper.BookResponse{
			Id:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Stock:    book.Stock,
			Price:    book.Price,
			Discount: book.Discount,
		})
	}

	return &tmp, nil
}

func (b *BookServiceImpl) DeleteBook(ctx context.Context, bookId string) error {
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)

	err := b.Repository.Delete(ctx, trx, bookId)

	if err != nil {
		return err
	}

	return nil
}

func (b *BookServiceImpl) GetBookById(ctx context.Context, bookId string) (*helper.BookResponse, error) {
	trx := b.DB.Begin()
	defer helper.CommitOrRollback(trx)

	book, err := b.Repository.FindById(ctx, trx, bookId)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, nil
	}

	return &helper.BookResponse{
		Id:       book.ID,
		Title:    book.Title,
		Author:   book.Author,
		Stock:    book.Stock,
		Price:    book.Price,
		Discount: book.Discount,
	}, nil

}
