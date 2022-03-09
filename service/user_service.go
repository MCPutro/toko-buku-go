package service

import (
	"context"
	"github.com/MCPutro/toko-buku-go/helper"
)

type UserService interface {
	CreateNewUser(ctx context.Context, user helper.UserCreateRequest) (*helper.UserCreateResponse, error)
	Login(ctx context.Context, user helper.UserLoginRequest) (*helper.UserLoginResponse, error)
	encodePassword(password string) (string, error)
	checkPassword(password, hash string) bool
}
