package api

import (
	"log"
	"net/http"
)

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
