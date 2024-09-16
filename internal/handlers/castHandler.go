package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "WatchifyAPI/internal/services"
	// "WatchifyAPI/internal/models" 
	// "log"
	// "WatchifyAPI/internal/db"
)

func GetCastByMovieId(w http.ResponseWriter, r *http.Request) {
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

    castMembers, err := services.GetCastByMovieId(movieId)
    if err != nil {
        http.Error(w, "Failed to retrieve crew members", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(castMembers)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(jsonData)
}
