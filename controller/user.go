package controller

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/denjos/curso/authentication"
	"github.com/denjos/curso/configuration"
	"github.com/denjos/curso/models"
	"log"
	"net/http"
)

// Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()
	c := sha256.Sum256([]byte(user.Password))
	//pwd:=base64.URLEncoding.EncodeToString(c[:32])
	pwd := fmt.Sprintf("%x", c)
	log.Println(pwd)
	db.Where("email=? and password=?", user.Email, pwd).First(&user)
	if user.ID > 0 {
		user.Password = ""
		token := authentication.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatal("error en parse token a json")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "usuario o clave incorrecto",
			Code:    http.StatusUnauthorized,
		}
		authentication.DisplayMessage(w, m)
	}
}

// UserCreate permite registrar un usuario
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("error de lectura de usuario: %s", err)
		m.Code = http.StatusBadRequest
		authentication.DisplayMessage(w, m)
		return
	}
	if user.Password != user.ConfirmPassword {
		m.Message = "los password no coinciden"
		m.Code = http.StatusBadRequest
		authentication.DisplayMessage(w, m)
		return
	}
	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	user.Picture = "https://gravatar.com/avatar/" + picstr + "?s=100"
	db := configuration.GetConnection()
	defer db.Close()
	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("error en la creacion del registro: %s", err)
		m.Code = http.StatusBadRequest
		authentication.DisplayMessage(w, m)
		return
	}
	m.Message = "Usuario creado con exito"
	m.Code = http.StatusOK
	authentication.DisplayMessage(w, m)

}
