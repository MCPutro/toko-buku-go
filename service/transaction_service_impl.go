package service

import (
	"context"
	"errors"
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

	//cek if stock >= quantity
	if book.Stock < request.Quantity {
		return nil, errors.New("stock not enough")
	}

	total := float32(request.Quantity) * book.Price
	discount := total * float32(book.Discount/100)
	amount := total - discount

	newTrx := &entity.Transaction{
		Date:     time.Now(),
		Customer: request.Customer,
		BookID:   request.BookID,
		Quantity: request.Quantity,
		Discount: book.Discount,
		Total:    amount,
	}

	trxId, err2 := t.TrxRepo.Save(ctx, t.db, newTrx)

	if err2 != nil {
		//update stock

	}

	newTrx.ID = trxId

	return newTrx, nil
}
