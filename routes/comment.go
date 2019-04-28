package routes

import (
	"github.com/denjos/curso/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller.CommentCreate).Methods("POST")
	subRouter.HandleFunc("/", controller.CommentGetAll).Methods("GET")
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controller.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)

}
