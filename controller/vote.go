package controller

import (
	"encoding/json"
	"errors"
	"github.com/denjos/curso/authentication"
	"github.com/denjos/curso/configuration"
	"github.com/denjos/curso/models"
	"net/http"
)

// VoteRegister controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}
	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = "error request incorrecto"
		authentication.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID
	db := configuration.GetConnection()
	defer db.Close()
	db.Where(" comment_id=? AND user_id=?", vote.CommentID, vote.UserID).Find(&currentVote)
	if currentVote.ID == 0 {
		db.Create(&vote)
		err = updateCommentVotes(vote.CommentID, vote.Value,false)
		if err != nil {
			m.Code = http.StatusNoContent
			m.Message = err.Error()
			authentication.DisplayMessage(w, m)
			return
		}
		m.Code = http.StatusCreated
		m.Message = "voto registrado"
		authentication.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value,true)
		if err != nil {
			m.Code = http.StatusBadRequest
			m.Message = err.Error()
			authentication.DisplayMessage(w, m)
			return
		}
		m.Code = http.StatusCreated
		m.Message = "voto actualizado"
		authentication.DisplayMessage(w, m)
		return
	}
	m.Code = http.StatusBadRequest
	m.Message = "este voto ya esta registrado"
	authentication.DisplayMessage(w, m)
}

// updateCommentVotes actualiza la cantidad de votos
func updateCommentVotes(commentID uint, vote bool,isUpdate bool) (err error) {
	comment := models.Comment{}
	db := configuration.GetConnection()
	defer db.Close()
	rows := db.First(&comment, commentID).RowsAffected
	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate {
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}
		db.Save(&comment)
	} else {
		err = errors.New("no se encontro un registro de comentario para asignarle un voto")
	}
	return
}
