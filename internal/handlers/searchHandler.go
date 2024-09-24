package handlers

import (
	"WatchifyAPI/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
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

	limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}

	for i, char := range value {
		if char == "+" {
			value[i] = " "
		}
	}

	movies, err := services.SearchByName(value[0], start, stop)
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

	limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}

	for i, char := range value {
		if char == "+" {
			value[i] = " "
		}
	}

	movies, err := services.SearchByActor(value[0], start, stop)
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
