package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(ctx context.Context, DB *gorm.DB, transaction *entity.Transaction) (uint8, error)
	FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.Transaction, error)
	FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.Transaction, error)
}
