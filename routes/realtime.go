package routes

import (
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
	"net/http"
)

// SetRealtimeRouter
func SetRealtimeRouter(router *mux.Router) {
	mel := melody.New()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})
	mel.HandleMessage(func(s *melody.Session, msg []byte) {
		mel.Broadcast(msg)
	})
}
