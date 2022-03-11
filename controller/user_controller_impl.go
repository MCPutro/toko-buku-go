package controller

import (
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
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

	helper.WriteToResponseBody2(w, err, newUserResponse)

	//var webResponse helper.Response
	//
	//if err != nil {
	//	webResponse = helper.Response{
	//		Status:  "error",
	//		Message: err.Error(),
	//		Data:    nil,
	//	}
	//} else {
	//	webResponse = helper.Response{
	//		Status: "success",
	//		Data:   newUserResponse,
	//	}
	//}
	//
	//helper.WriteToResponseBody(w, webResponse)
}

func (u *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	var userLogin helper.UserLoginRequest

	//identify content-type
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			helper.WriteToResponseBody(w, err)
			return
		}

		userLogin = helper.UserLoginRequest{
			Email:    r.PostFormValue("Email"),
			Password: r.PostFormValue("Password"),
		}
	} else {
		helper.ReadFromRequestBody(r, &userLogin)
	}

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

	if contentType == "application/x-www-form-urlencoded" {
		if fmt.Sprint(loginResponse.UserType) == fmt.Sprint(entity.Admin) {
			//redirect to list book admin
			http.Redirect(w, r, "/listBookAdmin", http.StatusSeeOther)
		} else {
			w.Write([]byte("bukan admin"))
		}
	} else {
		helper.WriteToResponseBody(w, webResponse)
	}

}
