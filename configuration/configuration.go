package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal("error archivo config.json")
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal("error parse configuration")
	}
	return c
}

//GetConnection conexion a la base de datos
func GetConnection() *gorm.DB {
	conf := getConfiguration()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Server, conf.Port, conf.User, conf.Password, conf.Database)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
