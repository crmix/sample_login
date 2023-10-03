package router

import (
	"login/api/controller"

	"github.com/gorilla/mux"
)

func Func(r *mux.Router, api controller.PostAPI){
	
	r.HandleFunc("/user",api.CreateUser).Methods("POST")
	r.HandleFunc("/login", api.Login).Methods("POST")
}