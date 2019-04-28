package routes

import (
	"github.com/denjos/curso/controller"
	"github.com/gorilla/mux"
)

func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controller.Login).Methods("POST")

}
