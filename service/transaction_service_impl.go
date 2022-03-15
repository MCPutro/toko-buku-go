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

func (t *TransactionServiceImpl) BuyBook(ctx context.Context, request helper.TransactionRequest) (*helper.TransactionResponse, error) {
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
	discount := total * (float32(book.Discount) / 100)
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

	return &helper.TransactionResponse{
		Id:       trxId,
		Date:     newTrx.Date,
		Customer: newTrx.Customer,
		BookID:   newTrx.BookID,
		Price:    book.Price,
		Quantity: newTrx.Quantity,
		Discount: newTrx.Discount,
		Total:    newTrx.Total,
	}, nil
}

func (t *TransactionServiceImpl) FindByCustomerEmail(ctx context.Context, email string) (*[]helper.TransactionResponse, error) {
	tx := t.db.Begin()
	defer helper.CommitOrRollback(tx)

	trxByCustomer, err := t.TrxRepo.FindByCustomer2(ctx, tx, email)

	if err != nil {
		return nil, err
	}

	var resp []helper.TransactionResponse

	for _, trx := range *trxByCustomer {
		resp = append(resp, helper.TransactionResponse{
			Id:        trx.Id,
			Date:      trx.Date,
			Customer:  trx.Customer,
			BookID:    trx.BookID,
			BookTitle: trx.BookTitle,
			Price:     trx.Price,
			Quantity:  trx.Quantity,
			Discount:  trx.Discount,
			Total:     trx.Total,
		})
	}

	if len(resp) == 0 {
		return nil, errors.New("data not found")
	} else {
		return &resp, nil
	}

}
