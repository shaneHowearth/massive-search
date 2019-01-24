package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	api "massive-search/grpc-api"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// RESTApp -
type RESTApp struct {
	Router *mux.Router
	GRPC   *grpc.ClientConn
}

// Initialise - Prepare router and routes for the application
func (a *RESTApp) Initialise() {
	a.Router = mux.NewRouter()
	a.initialiseRoutes()
	var err error
	a.GRPC, err = grpc.Dial("grpcservice:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to GRPC service with error %s", err)
	}
}

// Run -
func (a *RESTApp) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// Initialise routes, connect routes to their functions
func (a *RESTApp) initialiseRoutes() {
	a.Router.HandleFunc("/search/{word:[a-zA-Z]+}", a.getWord).Methods("GET")
}

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

func (a *RESTApp) getWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := api.NewStoredWordsClient(a.GRPC)
	response, err := c.GetWord(context.Background(), &api.Words{Word: []string{vars["word"]}})
	log.Printf("Response %+v", response)
	if err != nil {
		// TODO: We don't want the user to see this error message!
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	if response.Count == int32(-1) {
		respondWithError(w, http.StatusNotFound, "Not found")
	} else {
		respondWithJSON(w, http.StatusOK, &response)
	}
}
