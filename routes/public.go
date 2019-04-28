package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetPublicRouter(router *mux.Router)  {
	router.Handle("/",http.FileServer(http.Dir("./public")))
}
