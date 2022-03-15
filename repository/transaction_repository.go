package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(ctx context.Context, DB *gorm.DB, transaction *entity.Transaction) (string, error)
	FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.Transaction, error)
	FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.Transaction, error)
	FindByCustomer(ctx context.Context, DB *gorm.DB, Customer string) (*[]entity.Transaction, error)
	FindByCustomer2(ctx context.Context, DB *gorm.DB, email string) (*[]helper.TransactionResponse, error)
}
