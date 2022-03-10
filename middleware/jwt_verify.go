package middleware

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"regexp"
	"strings"
)

type JwtGate struct {
	service service.JwtService
}

func NewJwtGate(service service.JwtService) *JwtGate {
	return &JwtGate{service: service}
}

func (j *JwtGate) JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		//json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if !strings.Contains(header, "Bearer ") {
			w.WriteHeader(http.StatusForbidden)
			//json.NewEncoder(w).Encode("Missing auth token")
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: "Invalid Authorization Type",
				Data:    "Authorization must be Bearer Type",
			})
			return
		}

		auth := strings.Split(header, " ")

		token, err := j.service.ValidateToken(auth[1])
		if err != nil {
			helper.WriteToResponseBody(w, helper.Response{
				Status:  "error",
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		//if valid
		if token.Valid {
			path := r.URL.Path
			claims := token.Claims.(jwt.MapClaims)

			//if path is /book, only admin can access
			regex := regexp.MustCompile("/book/*")

			if regex.MatchString(path) {
				if claims["Role"].(float64) != 0 {
					w.WriteHeader(http.StatusForbidden)
					helper.WriteToResponseBody(w, helper.Response{
						Status:  "error",
						Message: "only admin can access",
					})
					return
				}
			} else {
				//if path isn't book
			}

		}
		next.ServeHTTP(w, r)
	})

}
