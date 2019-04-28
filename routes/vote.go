package routes

import (
	"github.com/denjos/curso/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetVoteRouter ruta para el registro de un voto
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller.VoteRegister).Methods("POST")
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controller.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
