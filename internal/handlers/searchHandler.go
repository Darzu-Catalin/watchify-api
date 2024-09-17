package handlers

import (
	"WatchifyAPI/internal/services"
	"encoding/json"
	"net/http"
	// "WatchifyAPI/internal/models"
	// "log"
	// "WatchifyAPI/internal/db"
)

func SearchMovieByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	value, ok := r.URL.Query()["title"]
	if !ok || len(value[0]) < 1 {
		http.Error(w, "Missing 'title' parameter", http.StatusBadRequest)
		return
	}

	for i, char := range value {
		if char == "+" {
			value[i] = " "
		}
	}

	movies, err := services.SearchByName(value[0])
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func SearchMovieByActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	value, ok := r.URL.Query()["name"]
	if !ok || len(value[0]) < 1 {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	for i, char := range value {
		if char == "+" {
			value[i] = " "
		}
	}

	movies, err := services.SearchByActor(value[0])
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
