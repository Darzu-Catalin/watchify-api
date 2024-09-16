package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "WatchifyAPI/internal/services"
	"WatchifyAPI/internal/models" 
	"log"
	"WatchifyAPI/internal/db"
)

func GetMovieById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    ids, ok := r.URL.Query()["id"]
    if !ok || len(ids[0]) < 1 {
        http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(ids[0])
    if err != nil {
        http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
        return
    }

    movie, err := services.GetMovieById(id)
    if err != nil {
        http.Error(w, "Movie not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(movie)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Parse the JSON body
    var movie models.Movie
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
        http.Error(w, "Invalid request payload"+err.Error(), http.StatusBadRequest)
        return
    }

    // Validate the movie data
    if movie.Title == ""  {
        http.Error(w, "Missing required movie fields", http.StatusBadRequest)
        return
    }

    // Add the movie to the database
    if err := services.AddMovie(movie); err != nil {
        http.Error(w, "Failed to add movie"+err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with a success message
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Movie added successfully"})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Extract movie ID from URL query parameters
    ids, ok := r.URL.Query()["id"]
    if !ok || len(ids[0]) < 1 {
        http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(ids[0])
    if err != nil {
        http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
        return
    }

    // Execute DELETE statement
    _, err = db.DB.Exec("DELETE FROM movies WHERE id = ?", id)
    if err != nil {
        log.Printf("Error deleting movie with ID %d: %v", id, err)
        http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusNoContent)
}

