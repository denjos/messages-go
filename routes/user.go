package routes

import (
	"github.com/denjos/curso/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetUserRouter ruta para el registro de usuario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller.UserCreate).Methods("POST")
	router.PathPrefix(prefix).Handler(negroni.New(negroni.Wrap(subRouter)))

}
