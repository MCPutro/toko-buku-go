package controller

import (
	"github.com/MCPutro/toko-buku-go/service"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func (u *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
