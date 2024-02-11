package main

import (
	"fmt"
	"os"

	"github.com/APoniatowski/code-task/pkg/api"
)

type envVars struct {
	hostname string
	database string
	username string
	password string
}

func dbConfig() *envVars {
	return &envVars{
		hostname: os.Getenv("HOSTNAME"),
		database: os.Getenv("POSTGRES_DB"),
		username: os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func main() {
	app := api.App{}
	fmt.Println("Initializing DB connection...")
	cfg := dbConfig()
	app.Initialize(cfg.hostname, "5432", cfg.database, cfg.username, cfg.password)
	fmt.Println("Starting API on port 8080...")
	defer app.DB.Close()
	go api.CleanupOldEntries(app.DB)
	app.Run(":8080")
}
