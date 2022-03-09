package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	Save(ctx context.Context, DB *gorm.DB, book entity.Book) (uint8, error)
	FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.Book, error)
	FindByTitle(ctx context.Context, DB *gorm.DB, title string) (*entity.Book, error)
	FindByTitleAndAuthor(ctx context.Context, DB *gorm.DB, title string, author string) (*entity.Book, error)
	FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.Book, error)
	Delete(ctx context.Context, DB *gorm.DB, id uint8) error
}
