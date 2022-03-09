package unit_test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type Season int64

const (
	Summer Season = iota
	Autumn
	Winter
	Spring
)

//func (s Season) String() string {
//	switch s {
//	case Summer:
//		return "summer"
//	case Autumn:
//		return "autumn"
//	case Winter:
//		return "winter"
//	case Spring:
//		return "spring"
//	}
//	return "unknown"
//}

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func TestName(t *testing.T) {
	printSeason(Winter)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TestPassword(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
