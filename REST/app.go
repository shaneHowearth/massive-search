package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RESTApp -
type RESTApp struct {
	Router *mux.Router
}

// Initialise - Prepare router and routes for the application
func (a *RESTApp) Initialise() {
	a.Router = mux.NewRouter()
	a.initialiseRoutes()
}

// Run -
func (a *RESTApp) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// Initialise routes, connect routes to their functions
func (a *RESTApp) initialiseRoutes() {}

// Send a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Printf("Error writing response %s", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
