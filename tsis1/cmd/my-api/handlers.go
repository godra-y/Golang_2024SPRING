package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	myapi "github.com/godra-y/tsis1/pkg/my-api"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, I'm Alnura. This is my Peaky Blinders characters API.")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Characters(w http.ResponseWriter, r *http.Request) {
	restaurants := myapi.GetСharacters()
	respondWithJSON(w, http.StatusOK, restaurants)
}

func Character(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	character, err := myapi.GetCharacter(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Character not found")
		return
	}

	respondWithJSON(w, http.StatusOK, character)
}
