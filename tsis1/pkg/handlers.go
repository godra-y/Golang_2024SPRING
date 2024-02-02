package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/godra-y/tsis1/api"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	log.Println("entering characters end point")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(api.Characters)
	if err != nil {
		//log.Printf("Failed to encode characters: %v", err)
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprintf(w, "Internal Server Error")
		return
	}

	w.Write(jsonResponse)
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	log.Println("entering characters end point")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	characterID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid character ID")
		return
	}

	for _, character := range api.Characters {
		if character.Id == characterID {
			json.NewEncoder(w).Encode(character)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Character not found")
}
