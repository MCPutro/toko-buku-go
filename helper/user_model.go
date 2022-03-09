package helper

import "github.com/MCPutro/toko-buku-go/entity"

type UserCreateRequest struct {
	Email    string
	UserName string
	Password string
	UserType entity.UserType
}

type UserCreateResponse struct {
	ID       uint8
	Email    string
	UserName string
	UserType entity.UserType
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type UserLoginResponse struct {
	ID       uint8
	UserName string
	Email    string
	Token    string
}
