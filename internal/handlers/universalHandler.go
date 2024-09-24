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

func GetMovieDataByMovieId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	movieId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	movieData, err := services.GetMovieDataByMovieId(movieId)
	if err != nil {
		http.Error(w, "Failed to retrieve crew members", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetNewMoviesByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
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

	movieData, err := services.GetNewMoviesByUserId(userId, start, stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetTrandingMoviesByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
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

	movieData, err := services.GetTrandingMoviesByUserId(userId, start, stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	genres, ok := r.URL.Query()["genre"]
	if !ok || len(genres[0]) < 1 {
		http.Error(w, "Missing 'genre' parameter", http.StatusBadRequest)
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


	movieData, err := services.GetMoviesByGenre(genres[0],start,stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetWatchLaterMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
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


	movieData, err := services.GetWatchLaterMovies(userId,start,stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}