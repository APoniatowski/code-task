package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (a *App) Initialize(dbHost, dbPort, dbName, dbUser, dbPassword string) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/messages", a.createMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/send", a.sendMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/{id}", a.deleteMessage).Methods("DELETE")
}
