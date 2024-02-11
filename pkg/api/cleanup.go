package api

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func CleanupOldEntries(db *sql.DB) {
	for {
		log.Println("Running cleanup...")
		fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

		_, err := db.Exec("DELETE FROM messages WHERE insert_time < $1", fiveMinutesAgo)
		if err != nil {
			log.Println("Error deleting old customer entries:", err)
		}

		time.Sleep(1 * time.Minute)
	}
}
