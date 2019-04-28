package authentication

import (
	"encoding/json"
	"github.com/denjos/curso/models"
	"github.com/labstack/gommon/log"
	"net/http"
)

func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal("error parse message a json : %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
