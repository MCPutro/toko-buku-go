package service

import (
	"context"
	"errors"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *gorm.DB
	Jwt      JwtService
}

func NewUserService(userRepo repository.UserRepository, DB *gorm.DB, jwt JwtService) UserService {
	return &UserServiceImpl{UserRepo: userRepo, DB: DB, Jwt: jwt}
}

func (u *UserServiceImpl) CreateNewUser(ctx context.Context, user helper.UserCreateRequest) (*helper.UserCreateResponse, error) {
	trx := u.DB.Begin()
	defer helper.CommitOrRollback(trx)

	encodePass, err := u.encodePassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: encodePass,
		UserType: user.UserType,
	}

	UserId, err := u.UserRepo.Save(ctx, trx, newUser)
	if err != nil {
		return nil, err
	}

	return &helper.UserCreateResponse{
		ID:       UserId,
		Email:    newUser.Email,
		UserName: newUser.UserName,
		UserType: newUser.UserType,
	}, nil
}

func (u *UserServiceImpl) Login(ctx context.Context, user helper.UserLoginRequest) (*helper.UserLoginResponse, error) {
	trx := u.DB.Begin()
	defer helper.CommitOrRollback(trx)

	existingUser, err := u.UserRepo.FindByEmail(ctx, trx, user.Email)
	if err != nil {
		return nil, err
	}

	//check pass is match
	checkPassword := u.checkPassword(user.Password, existingUser.Password)

	if !checkPassword {
		return nil, errors.New("invalid credential")
	} else {
		token := u.Jwt.GenerateToken(existingUser.Email, existingUser.UserType)

		return &helper.UserLoginResponse{
			ID:       existingUser.ID,
			UserName: existingUser.UserName,
			Email:    existingUser.Email,
			UserType: existingUser.UserType,
			Token:    token,
		}, nil
	}
}

func (u *UserServiceImpl) encodePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserServiceImpl) checkPassword(password string, encodePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password))
	return err == nil
}
