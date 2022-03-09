package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, DB *gorm.DB, user entity.User) (uint8, error)
	FindByEmail(ctx context.Context, DB *gorm.DB, email string) (*entity.User, error)
	FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.User, error)
}
