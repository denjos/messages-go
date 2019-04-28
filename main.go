package main

import (
	"flag"
	"fmt"
	"github.com/denjos/curso/commons"
	"github.com/denjos/curso/migration"
	"github.com/denjos/curso/routes"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "genera el migration to database")
	flag.IntVar(&commons.Port, "port", 8080, "start web server port")
	flag.Parse()
	if migrate == "yes" {
		log.Println("start migration")
		migration.Migrate()
		log.Println("finish migration")
	}
	router := routes.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	// Inicio del servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d",commons.Port),
		Handler: n,
	}
	log.Printf("starting web server at http://localhost:%d",commons.Port)
	err := server.ListenAndServe()
	log.Println("error en el inicio del servidor", err)
}
