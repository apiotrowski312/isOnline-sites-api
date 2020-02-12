package app

import (
	"log"
	"net/http"

	"github.com/apiotrowski312/isOnline-sites-api/src/controllers/sites"
	"github.com/gorilla/mux"
)

func StartApplication() {
	r := mux.NewRouter()

	r.HandleFunc("/sites/{id}", sites.Get).Methods("GET")
	r.HandleFunc("/sites", sites.GetUserSites).Methods("GET")
	r.HandleFunc("/sites", sites.Post).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", r))
}
