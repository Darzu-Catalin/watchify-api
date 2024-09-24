package handlers

import (
    "encoding/json"
    "net/http"
    "WatchifyAPI/internal/services"
    "WatchifyAPI/internal/models"
    "strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
    var user models.User

    // Decode the request body into the User struct
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call the service layer to add the user
    createdUser, err := services.CreateUser(user)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Return the created user as the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdUser)
}

func AddUserInteraction(w http.ResponseWriter, r *http.Request) {
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

    movie, ok := r.URL.Query()["movieId"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'movie id' parameter", http.StatusBadRequest)
		return
	}
    
	movieId, err := strconv.Atoi(movie[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}


    interaction, ok := r.URL.Query()["interaction"]
	if !ok || len(interaction[0]) < 1 {
		http.Error(w, "Missing 'interaction' parameter", http.StatusBadRequest)
		return
	}
    // Decode the request body into the User struct
    // Call the service layer to add the user
    response, err := services.AddUserInteraction(userId,movieId,interaction[0])
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Return the created user as the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}



