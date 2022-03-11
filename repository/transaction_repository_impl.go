package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (t *TransactionRepositoryImpl) Save(ctx context.Context, DB *gorm.DB, transaction *entity.Transaction) (uint8, error) {
	result := DB.WithContext(ctx).Create(&transaction)

	if result.Error != nil {
		return 0, result.Error
	}

	return transaction.ID, nil

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
		return nil, nil
	}

}
