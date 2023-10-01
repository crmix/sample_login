package controller

import (
	//"fmt"
	"log"
	"login/service"
	"net/http"
	"login/model"

)

type userApi struct {
	service service.UserService
}

func NewUserAPI(service service.UserService) UserAPI {
	return &userApi{
		service: service,
	}
}

type UserAPI interface {
	GetUsersName(w http.ResponseWriter, r *http.Request)
}

func (api *userApi) GetUsersName(w http.ResponseWriter, r *http.Request) {

	res, err := api.service.GetUsersName([]model.User{})
	if err != nil {
		log.Println(err)
		return
	}

	WriteSuccess(w, res)
}
