package controller

import (
	"encoding/json"
	"fmt"
	"github.com/denjos/curso/authentication"
	"github.com/denjos/curso/commons"
	"github.com/denjos/curso/configuration"
	"github.com/denjos/curso/models"
	"log"
	"net/http"
	"strconv"
	"github.com/olahol/melody"
	"golang.org/x/net/websocket"
)

// Melody use realtime
var Melody *melody.Melody

func init()  {
	Melody =melody.New()
}

func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	user := models.User{}
	user, _ = r.Context().Value("user").(models.User)
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("request invalido %s", err)
		authentication.DisplayMessage(w, m)
		return
	}
	comment.UserID = user.ID
	db := configuration.GetConnection()
	defer db.Close()
	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("error en la creacion del registro: %s", err)
		authentication.DisplayMessage(w, m)
		return
	}

	db.Model(&comment).Related(&comment.User)
	comment.User[0].Password=""
	j,err:=json.Marshal(&comment)
	if err != nil {
		m.Message=fmt.Sprintf("error parse comment o json: %s",err)
		m.Code=http.StatusInternalServerError
		authentication.DisplayMessage(w,m)
		return
	}
	origin:=fmt.Sprintf("http://localhost:%d/",commons.Port)
	url:=fmt.Sprintf("ws://localhost:%d/ws",commons.Port)
	ws,err:=websocket.Dial(url,"",origin)
	if err != nil {
		log.Fatal(err)
	}

	if _,err:=ws.Write(j);err!=nil {
		log.Fatal(err)
	}
	
	m.Code = http.StatusCreated
	m.Message = "comentario creado con exito"
	authentication.DisplayMessage(w, m)
}

// CommentGetAll obtiene todos los mensajes
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}

	user, _ = r.Context().Value("user").(models.User)
	//http://sufuq.com/books/golang/The%20Way%20To%20Go.pdf
	//https://github.com/adonovan/gopl.io/blob/master/ch1/dup1/main.go
	//https://github.com/dariubs/GoBooks
	vars := r.URL.Query()
	db := configuration.GetConnection()
	defer db.Close()
	cComment := db.Where("parent_id=0")
	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			cComment = cComment.Order("votes desc,created_at")

		}
	} else {
		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("error:", err)
			}
			cComment = cComment.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}
		cComment = cComment.Order("id desc")
	}
	cComment.Find(&comments)

	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User)
		comments[i].User[0].Password = ""
		comments[i].Children = commentGetChildren(comments[i].ID)
		// Se busca el voto del usuario en sesion
		vote.CommentID = comments[i].ID
		vote.UserID = user.ID
		count := db.Where(&vote).Find(&vote).RowsAffected
		if count > 0 {
			if vote.Value {
				comments[i].HasVote = 1
			} else {
				comments[i].HasVote = -1
			}
		}
	}

	j, err := json.Marshal(comments)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "error el parse de los comentarios"
		authentication.DisplayMessage(w, m)
		return
	}
	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.Code = http.StatusNoContent
		m.Message = "no se encontraron comentarios"
		authentication.DisplayMessage(w, m)
	}
}

func commentGetChildren(id uint) (children []models.Comment) {
	db := configuration.GetConnection()
	defer db.Close()
	db.Where("parent_id = ?", id).Find(&children)
	for i := range children {
		db.Model(&children[i]).Related(&children[i].User)
		children[i].User[0].Password = ""
	}
	return
}
