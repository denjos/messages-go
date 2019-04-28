package migration

import (
	"github.com/denjos/curso/configuration"
	"github.com/denjos/curso/models"
)

//Migrate permite crear las tablas en la base de datos
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()
	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")

}
