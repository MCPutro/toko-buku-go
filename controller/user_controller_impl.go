package controller

import (
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{service: service}
}

func (u *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser helper.UserCreateRequest
	helper.ReadFromRequestBody(r, &newUser)

	newUserResponse, err := u.service.CreateNewUser(r.Context(), newUser)

	var webResponse helper.Response

	if err != nil {
		webResponse = helper.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		webResponse = helper.Response{
			Status: "success",
			Data:   newUserResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (u *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	var userLogin helper.UserLoginRequest
	helper.ReadFromRequestBody(r, &userLogin)

	loginResponse, err := u.service.Login(r.Context(), userLogin)

	var webResponse helper.Response

	if err != nil {
		webResponse = helper.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
	} else {
		webResponse = helper.Response{
			Status:  "success",
			Message: "",
			Data:    loginResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}
