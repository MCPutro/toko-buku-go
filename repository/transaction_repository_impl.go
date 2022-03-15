package repository

import (
	"context"
	"errors"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (t *TransactionRepositoryImpl) Save(ctx context.Context, DB *gorm.DB, transaction *entity.Transaction) (string, error) {
	result := DB.WithContext(ctx).Create(&transaction)

	if result.Error != nil {
		return "", result.Error
	}

	return transaction.Id, nil

}

func (t *TransactionRepositoryImpl) FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.Transaction, error) {
	var tmp entity.Transaction

	existingTrx := DB.WithContext(ctx).First(&tmp, id)

	if existingTrx.Error != nil {
		return nil, existingTrx.Error
	}

	if existingTrx.RowsAffected > 0 {
		return &tmp, nil
	} else {
		return nil, nil
	}
}

func (t *TransactionRepositoryImpl) FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.Transaction, error) {
	var tmp []entity.Transaction

	resultFindAll := DB.WithContext(ctx).Find(&tmp)

	if resultFindAll.Error != nil {
		return nil, resultFindAll.Error
	}

	if resultFindAll.RowsAffected > 0 {
		return &tmp, nil
	} else {
		return nil, errors.New("no data transaction")
	}
}

func (t *TransactionRepositoryImpl) FindByCustomer(ctx context.Context, DB *gorm.DB, Customer string) (*[]entity.Transaction, error) {
	var tmp []entity.Transaction

	trxByCustomer := DB.WithContext(ctx).Where("customer = ?", Customer).Find(&tmp)

	if trxByCustomer.Error != nil {
		return nil, trxByCustomer.Error
	}

	if trxByCustomer.RowsAffected > 0 {
		return &tmp, nil
	} else {
		return nil, errors.New("no data transaction")
	}
}

func (t *TransactionRepositoryImpl) FindByCustomer2(ctx context.Context, DB *gorm.DB, email string) (*[]helper.TransactionResponse, error) {

	var transactions []helper.TransactionResponse

	err := DB.WithContext(ctx).Raw("select  t.id, t.\"date\", t.customer, t.book_id, b.title as book_title, b.price, t.quantity, t.discount, t.total  from  transactions t, books b where t.book_id = b.id and t.customer = ?", email).Scan(&transactions).Error

	if err != nil {
		return nil, err
	}

	return &transactions, nil

	//rows, err := DB.WithContext(ctx).
	//	Raw("select  t.id, t.\"date\", t.customer, t.book_id, b.title as book_title, b.price, t.quantity, t.discount, t.total  from  transactions t, books b where t.book_id = b.id and t.customer = ?", email).
	//	Rows()
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer rows.Close()
	//
	//var transaction helper.TransactionResponse
	//for rows.Next() {
	//	err2 := rows.Scan(&transaction.Id, &transaction.Date, &transaction.Customer, &transaction.BookID, &transaction.BookTitle, &transaction.Price, &transaction.Quantity, &transaction.Discount, &transaction.Total)
	//	if err2 != nil {
	//		fmt.Println("error print :", err2)
	//	}
	//	fmt.Println(transaction)
	//	transactions = append(transactions, transaction)
	//}
	//
	//fmt.Println("---")
	//fmt.Println(transactions)
	//
	//return &transactions, nil
}
