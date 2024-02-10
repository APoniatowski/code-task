package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
