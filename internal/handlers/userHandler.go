package handlers

import (
    "encoding/json"
    "net/http"
    "WatchifyAPI/internal/services"
    "WatchifyAPI/internal/models"
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
