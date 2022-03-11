package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (b *BookRepositoryImpl) Save(ctx context.Context, DB *gorm.DB, book *entity.Book) (uint8, error) {
	var result *gorm.DB

	if DB.WithContext(ctx).Where("id = ?", book.ID).Updates(&book).RowsAffected == 0 {
		result = DB.WithContext(ctx).Create(&book)
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return book.ID, nil
}

func (b *BookRepositoryImpl) FindAll(ctx context.Context, DB *gorm.DB) (*[]entity.Book, error) {
	var listBook []entity.Book

	find := DB.WithContext(ctx).Find(&listBook)

	if find.Error != nil {
		return nil, find.Error
	}

	if find.RowsAffected > 0 {
		return &listBook, nil
	} else {
		return nil, nil
	}

}

func (b *BookRepositoryImpl) FindByTitle(ctx context.Context, DB *gorm.DB, title string) (*entity.Book, error) {
	var book entity.Book

	find := DB.WithContext(ctx).Where("title = ?", title).Find(&book)

	if find.Error != nil {
		return nil, find.Error
	}

	if find.RowsAffected > 0 {
		return &book, nil
	} else {
		return nil, nil
	}

}

func (b *BookRepositoryImpl) FindByTitleAndAuthor(ctx context.Context, DB *gorm.DB, title string, author string) (*entity.Book, error) {
	var book entity.Book

	find := DB.WithContext(ctx).Where("title = ? AND author = ?", title, author).First(&book)

	if find.Error != nil {
		return nil, find.Error
	}

	if find.RowsAffected > 0 {
		return &book, nil
	} else {
		return nil, nil
	}

}

func (b *BookRepositoryImpl) FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.Book, error) {
	var book entity.Book

	find := DB.WithContext(ctx).First(&book, id)

	if find.Error != nil {
		return nil, find.Error
	}

	if find.RowsAffected > 0 {
		return &book, nil
	} else {
		return nil, nil
	}
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, id uint8) error {
	tx := DB.WithContext(ctx).Delete(&entity.Book{}, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
