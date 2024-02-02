package main

import (
	"log"
	"net/http"

	handlers "github.com/godra-y/tsis1/pkg"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routers")

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/characters", handlers.GetCharacters).Methods("GET")
	router.HandleFunc("/characters/{id}", handlers.GetCharacter).Methods("GET")

	http.Handle("/", router)
	log.Println("API server is ready to serve")
	http.ListenAndServe(":8080", router)
}
