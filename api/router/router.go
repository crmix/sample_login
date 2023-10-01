package router

import(
	"login/api/controller"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, api controller.UserAPI){
	r.HandleFunc("/user", api.GetUsersName).Methods("GET")
	//r.HandleFunc("/user", api.PostNewUser).Methods("POST")
}

func Func(r *mux.Router, api controller.PostAPI){
	r.HandleFunc("/user",api.PostNewUser).Methods("POST")
}