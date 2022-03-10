package service

import (
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/golang-jwt/jwt/v4"
)

type JwtService interface {
	GenerateToken(id string, userType entity.UserType) string
	ValidateToken(token string) (*jwt.Token, error)
}
