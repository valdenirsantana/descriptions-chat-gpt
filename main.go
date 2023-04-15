package main

import (
	"log"
	"net/http"

	"chat-gpt/internal/descriptions/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/descriptions/{productName}", handlers.GetDescriptions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
