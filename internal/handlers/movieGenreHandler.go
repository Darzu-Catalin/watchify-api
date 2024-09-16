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

func GetGenreByMovieId(w http.ResponseWriter, r *http.Request) {
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

    movieGenre, err := services.GetGenreByMovieId(movieId)
    if err != nil {
        http.Error(w, "Failed to retrieve crew members", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(movieGenre)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}

func GetAllGenres(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var genres []string
    var err error
    genres, err = services.GetAllGenres();
    if err != nil {
        http.Error(w, "Failed to retrieve all genres", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(genres)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}

