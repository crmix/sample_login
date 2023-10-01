package router

import(
	"login/api/controller"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, api controller.UserAPI){
	r.HandleFunc("/user", api.GetUsersName).Methods("GET")
}