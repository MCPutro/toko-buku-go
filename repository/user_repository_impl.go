package repository

import (
	"context"
	"github.com/MCPutro/toko-buku-go/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, DB *gorm.DB, user entity.User) (uint8, error) {
	result := DB.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, DB *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	tx := DB.WithContext(ctx).Where("email = ?", email).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, DB *gorm.DB, id uint8) (*entity.User, error) {
	var user entity.User
	tx := DB.WithContext(ctx).First(&user, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
