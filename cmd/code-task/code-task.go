package main

import (
	"fmt"

	"github.com/APoniatowski/code-task/pkg/api"
)

func main() {
	app := api.App{}
	fmt.Println("Initializing DB connection...")
	app.Initialize("localhost", "5432", "somedb", "somdbuser", "somedbpassword")
	fmt.Println("Starting API on port 8080...")
	app.Run(":8080")
}
