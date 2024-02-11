package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) getMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming GET (getMessages) request from ", r.Header.Get("X-Real-IP"))
	rows, err := a.DB.Query("SELECT id, email, title, content, mailing_id, insert_time FROM messages")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Email, &msg.Title, &msg.Content, &msg.MailingID, &msg.InsertTime)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		messages = append(messages, msg)
	}
	if err = rows.Err(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

func (a *App) createMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming POST (createMessage) request from ", r.Header.Get("X-Real-IP"))
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println("Error decoding JSON payload:", err)
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
	fmt.Println("Incoming POST (sendMessage) request from ", r.Header.Get("X-Real-IP"))
	// TODO: Send message (mocked)

	_, err := a.DB.Exec("DELETE FROM messages WHERE mailing_id = $1", 1) // Mocked mailing_id
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Messages sent and deleted"})
}

func (a *App) deleteMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming DELETE (deleteMessage) request from ", r.Header.Get("X-Real-IP"))
	params := mux.Vars(r)
	id := params["id"]

	_, err := a.DB.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Message deleted"})
}
