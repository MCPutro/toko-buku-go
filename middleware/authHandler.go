package middleware

import (
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type AuthCheck struct {
	service service.JwtService
}

func NewAuthCheck(service service.JwtService) *AuthCheck {
	return &AuthCheck{service: service}
}

func (c *AuthCheck) AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string

		secChUa := r.Header.Get("Sec-Ch-Ua")

		if secChUa != "" {
			tokenString = "Bearer " + helper.GetCookie(r, "token")
		} else {
			tokenString = r.Header.Get("Authorization")
		}

		if !strings.Contains(tokenString, "Bearer ") {
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: "Invalid Authorization Type",
			})
			return
		}

		auth := strings.Split(tokenString, " ")

		token, err := c.service.ValidateToken(auth[1])
		if err != nil {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if fmt.Sprint(claims["Role"]) != fmt.Sprint(entity.Admin) {
			w.WriteHeader(http.StatusForbidden)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: "only admin can access",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *AuthCheck) AuthCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string

		secChUa := r.Header.Get("Sec-Ch-Ua")

		if secChUa != "" {
			tokenString = "Bearer " + helper.GetCookie(r, "token")
		} else {
			tokenString = r.Header.Get("Authorization")
		}

		if !strings.Contains(tokenString, "Bearer ") {
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: "Invalid Authorization Type",
			})
			return
		}

		auth := strings.Split(tokenString, " ")

		token, err := c.service.ValidateToken(auth[1])
		if err != nil {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: "internal server error",
			})
			return
		}

	})
}
