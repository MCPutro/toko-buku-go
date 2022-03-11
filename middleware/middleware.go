package middleware

import (
	"fmt"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/gorilla/mux"
	"net/http"
)

var listExcludeJWT = map[string]bool{
	"/user/SignIn": true,
	"/user/SignUp": true,
	"/":            true,
}

type Middleware struct {
	Route   *mux.Router
	service service.JwtService
}

func NewMiddleware(route *mux.Router, service service.JwtService) *Middleware {
	return &Middleware{Route: route, service: service}
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//var header = r.Header.Get("Authorization")
	path := r.URL.Path

	fmt.Println("middleware:", path)

	//
	////print result
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	//if !listExcludeJWT[path] {
	//	if !strings.Contains(header, "Bearer ") {
	//		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
	//		helper.WriteToResponseBody(w, helper.Response{
	//			Status:  "error",
	//			Message: "Invalid Authorization Type",
	//			Data:    "Authorization must be Bearer Type",
	//		})
	//		return
	//	}
	//
	//	auth := strings.Split(header, " ")
	//
	//	token, err := m.service.ValidateToken(auth[1])
	//	if err != nil {
	//		helper.WriteToResponseBody(w, helper.Response{
	//			Status:  "error",
	//			Message: err.Error(),
	//			Data:    nil,
	//		})
	//		return
	//	}
	//
	//	claims := token.Claims.(jwt.MapClaims)
	//	//if path is /book, only admin can access
	//	regex := regexp.MustCompile("/book/*")
	//
	//	if regex.MatchString(path) {
	//		if fmt.Sprint(claims["Role"]) != fmt.Sprint(entity.Admin) {
	//			w.WriteHeader(http.StatusForbidden)
	//			helper.WriteToResponseBody(w, helper.Response{
	//				Status:  "error",
	//				Message: "only admin can access",
	//			})
	//			return
	//		}
	//	}
	//}

	m.Route.ServeHTTP(w, r)

}
