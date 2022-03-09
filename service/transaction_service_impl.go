package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"gorm.io/gorm"
	"time"
)

type TransactionServiceImpl struct {
	TrxRepo  repository.TransactionRepository
	BookRepo repository.BookRepository
	//UserRepo repository.UserRepository
	db *gorm.DB
}

func NewTransactionService(trxRepo repository.TransactionRepository, bookRepo repository.BookRepository, db *gorm.DB) TransactionService {
	return &TransactionServiceImpl{TrxRepo: trxRepo, BookRepo: bookRepo, db: db}
}

func (t *TransactionServiceImpl) BuyBook(ctx context.Context, request helper.TransactionRequest) (*entity.Transaction, error) {
	book, err := t.BookRepo.FindById(ctx, t.db, request.BookID)
	if err != nil {
		return nil, errors.New("book not found")
	}

	fmt.Println(book)

	total := request.Quantity * book.Price
	discount := total * book.Discount
	amount := total - discount

	fmt.Println(amount)

	newTrx := entity.Transaction{
		Date:     time.Now(),
		Customer: request.Customer,
		BookID:   request.BookID,
		Quantity: request.Quantity,
		Discount: book.Discount,
		Total:    amount,
	}

	trxId, err := t.TrxRepo.Save(ctx, t.db, newTrx)

	newTrx.ID = trxId

	return &newTrx, nil
}
