package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/google/uuid"
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
	tx := t.db.Begin()
	defer helper.CommitOrRollback(tx)

	book, err := t.BookRepo.FindById(ctx, tx, request.BookID)
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
		Id:       uuid.New().String(),
		Date:     time.Now(),
		Customer: request.Customer,
		BookID:   request.BookID,
		Quantity: request.Quantity,
		Discount: book.Discount,
		Total:    amount,
	}

	trxId, errCreateTrx := t.TrxRepo.Save(ctx, tx, newTrx)

	if errCreateTrx != nil {
		panic(errCreateTrx)

		return nil, errCreateTrx
	} else {
		//update stock
		newStock := book.Stock - newTrx.Quantity
		errUpdateStock := t.BookRepo.UpdateStock(ctx, tx, book.ID, newStock)
		if errUpdateStock != nil {
			fmt.Println(errUpdateStock)
			panic(errUpdateStock)
		}
	}

	newTrx.Id = trxId

	return newTrx, nil
}
