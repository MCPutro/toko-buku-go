package service

import "github.com/golang-jwt/jwt/v4"

type JwtService interface {
	GenerateToken(id string) string
	ValidateToken(token string) (*jwt.Token, error)
}
