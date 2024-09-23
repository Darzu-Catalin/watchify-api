package services

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/models"
	"log"
)

func CreateUser(user models.User) (*models.User, error) {
	query := "INSERT INTO users (id, email, name, profile_picture) VALUES (?,?, ?, ?)"


	// Execute the SQL query
	res, err := db.DB.Exec(query, user.ID, user.Email, user.Nickname, user.Profile_picture)
	if err != nil {
		log.Printf("Error executing SQL query: %v", err) // Log the error
		return nil, err
	}

	// Retrieve the last inserted ID
	userId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error retrieving last inserted ID: %v", err) // Log the error
		return nil, err
	}
	id := int(userId)
	user.ID = &id

	// Return the created user
	return &user, nil
}
