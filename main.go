package main

import (
	"log"
	"net/http"
	"login/api/controller"
	"login/api/router"
	"login/config"
	"login/platform"
	"login/service"
	"login/storage"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)


func main() {
	db := platform.DBConn()
	cfg := config.Load()

    storage :=storage.PostRepo(db)
	service :=service.PostUserService(storage, *zap.NewStdLog(zap.L()))
	api := controller.PostUserAPI(service)
	root := mux.NewRouter()

	router.Func(root, api)

	httpServer := http.Server{
		Addr: cfg.HTTPPort,
		Handler: root,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("shut down server")
	}

}
