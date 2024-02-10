package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Message struct {
	Email      string    `json:"email"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	MailingID  int       `json:"mailing_id"`
	InsertTime time.Time `json:"insert_time"`
}

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

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

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/messages", a.createMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/send", a.sendMessage).Methods("POST")
	a.Router.HandleFunc("/api/messages/{id}", a.deleteMessage).Methods("DELETE")
}

func (a *App) createMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	_, err = a.DB.Exec("INSERT INTO messages (email, title, content, mailing_id, insert_time) VALUES ($1, $2, $3, $4, $5)",
		msg.Email, msg.Title, msg.Content, msg.MailingID, msg.InsertTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, msg)
}

func (a *App) sendMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Send message (mocked)

	_, err := a.DB.Exec("DELETE FROM messages WHERE mailing_id = $1", 1) // Mocked mailing_id
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Messages sent and deleted"})
}

func (a *App) deleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := a.DB.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Message deleted"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	app := App{}
	fmt.Println("Initializing DB connection...")
	app.Initialize("localhost", "5432", "somedb", "somdbuser", "somedbpassword")
	fmt.Println("Starting API on port 8080...")
	app.Run(":8080")
}
