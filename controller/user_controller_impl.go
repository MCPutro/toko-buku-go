package controller

import (
	"fmt"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/service"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{service: service}
}

func (u *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser helper.UserCreateRequest

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			helper.WriteToResponseBody(w, err)
			return
		}

		i, _ := strconv.ParseInt(r.PostFormValue("UserType"), 10, 8)

		newUser = helper.UserCreateRequest{
			Email:    r.PostFormValue("Email"),
			UserName: r.PostFormValue("Username"),
			Password: r.PostFormValue("Password"),
			UserType: entity.UserType(i),
		}
	} else {
		helper.ReadFromRequestBody(r, &newUser)
	}

	newUserResponse, err := u.service.CreateNewUser(r.Context(), newUser)

	if contentType == "application/x-www-form-urlencoded" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		helper.WriteToResponseBody2(w, err, newUserResponse)
	}

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

	if contentType == "application/x-www-form-urlencoded" {
		helper.SetCookie(w, "email", userLogin.Email, 60)
		helper.SetCookie(w, "token", loginResponse.Token, 60)
		if fmt.Sprint(loginResponse.UserType) == fmt.Sprint(entity.Admin) {
			//redirect to list book admin
			http.Redirect(w, r, "/listBookAdmin", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/listBook", http.StatusSeeOther)
		}
	} else {
		helper.WriteToResponseBody2(w, err, loginResponse)
	}

}
