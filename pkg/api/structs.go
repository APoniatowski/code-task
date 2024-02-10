package api

import (
	"database/sql"
	"time"

	"github.com/gorilla/mux"
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
