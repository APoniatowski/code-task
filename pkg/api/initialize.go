package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (a *App) Initialize(dbHost, dbPort, dbName, dbUser, dbPassword string) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Printf("Failed to connect to the database: %v. Retrying...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping the database: %v. Retrying...", err)
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database after retries: %v", err)
	}

	a.DB = db

	err = a.createTables()
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) createTables() error {
	fmt.Printf("Creating tables... ")
	_, err := a.DB.Exec(`
        CREATE TABLE IF NOT EXISTS messages (
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            mailing_id INT NOT NULL,
            insert_time TIMESTAMP NOT NULL
        );
    `)
	if err != nil {
		return err
	}
	fmt.Println("Done.")
	return nil
}

func (a *App) initializeRoutes() {
	fmt.Printf("Initializing routes... ")
	a.Router.HandleFunc("/api/messages", a.getMessages).Methods("GET")
	a.Router.HandleFunc("/api/messages", a.createMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/send", a.sendMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/{id}", a.deleteMessage).Methods("DELETE")
	fmt.Println("Initialized.")
}
