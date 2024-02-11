package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func CleanupOldEntries(db *sql.DB) {
	fmt.Println("Starting Cleaner...")
	for {
		oldMessages, err := getOldMessages(db)
		if err != nil {
			log.Println("Error getting old customer entries:", err)
		}
		if len(oldMessages) > 0 {
			_, err := db.Exec("DELETE FROM messages WHERE insert_time < NOW() - INTERVAL '5 minutes';")
			if err != nil {
				log.Println("Error deleting old customer entries:", err)
			}
			log.Println("Running cleanup...")

		}
		time.Sleep(1 * time.Minute)
	}
}

func getOldMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query("SELECT * FROM messages WHERE insert_time < NOW() - INTERVAL '5 minutes';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Email, &msg.Title, &msg.Content, &msg.MailingID, &msg.InsertTime)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
