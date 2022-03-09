package main

import (
	"github.com/MCPutro/toko-buku-go/config"
	"github.com/MCPutro/toko-buku-go/controller"
	"github.com/MCPutro/toko-buku-go/entity"
	"github.com/MCPutro/toko-buku-go/helper"
	"github.com/MCPutro/toko-buku-go/repository"
	"github.com/MCPutro/toko-buku-go/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := config.GetConnection()
	err := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		return
	}

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	r := mux.NewRouter()

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/SignUp", userController.SignUp)
	user.HandleFunc("/SignIn", userController.SignIn)

	err2 := http.ListenAndServe(":8080", r)
	if err2 != nil {
		helper.PanicIfError(err2)
	}
}
