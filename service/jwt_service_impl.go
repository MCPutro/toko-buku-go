package service

import (
	"errors"
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type jwtServiceImpl struct {
	secretKey string
	issuer    string
}

func NewJwtServiceImpl() JwtService {
	return &jwtServiceImpl{
		secretKey: getSecretKey(),
		issuer:    "MCPutro",
	}
}

type jwtCustomClaim struct {
	Id   string
	Role entity.UserType
	jwt.RegisteredClaims
}

func (j *jwtServiceImpl) GenerateToken(id string, userType entity.UserType) string {
	claims := jwtCustomClaim{
		Id:   id,
		Role: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.secretKey))
	//fmt.Printf("%v %v", ss, err)
	if err != nil {
		return ""
	}
	return ss
}

func (j *jwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if t.Valid {
		//fmt.Println("You look nice today")
		return t, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		//fmt.Println("That's not even a token")
		return nil, errors.New("that's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		//fmt.Println("Timing is everything")
		return nil, errors.New("token is expired")
	} else {
		//fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}

func getSecretKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return ""
	}

	secretKey := os.Getenv("KEY_SECRET")
	if secretKey == "" {
		secretKey = "system"
	}

	return secretKey
}
