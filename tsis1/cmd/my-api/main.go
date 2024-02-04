package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routers")

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/characters", Characters).Methods("GET")
	router.HandleFunc("/characters/{id}", Character).Methods("GET")

	http.Handle("/", router)
	log.Println("API server is ready to serve")
	http.ListenAndServe(":8080", router)
}
