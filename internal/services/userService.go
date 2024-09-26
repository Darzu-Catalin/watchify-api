package services

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/models"
	"log"
)

func CreateUser(user models.User) (*models.User, error) {
	query := "INSERT INTO users (id, email, name) VALUES (?,?, ?)"


	// Execute the SQL query
	res, err := db.DB.Exec(query, user.ID, user.Email, user.Nickname)
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

func AddUserInteraction(id int, movieId int, interaction string) (string, error) {
	query := "INSERT INTO user_movie_interactions (user_id, movie_id, interaction_type) VALUES (?,?, ?)"
	var response string
	
	// Execute the SQL query
	res, err := db.DB.Exec(query, id, movieId, interaction)
	if err != nil {
		log.Printf("Error executing SQL query: %v", err)
		response = "error" // Log the error
		return response, err
	}

	rowsAffected, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error retrieving rows affected: %v", err)
		response = "error"
        return response, err
    }
	response = "Succesfully added interaction"
    log.Printf("Rows affected: %d", rowsAffected)

	// Return the created user
	return response, err
}

func GetAllUserInteraction(id int, movieId int) ([]string, error) {
	var interactions []string
	query := "SELECT interaction_type FROM user_movie_interactions WHERE user_id = ? and movie_id = ?"
	
	
	rows, err := db.DB.Query(query, id, movieId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var interaction string
		if err := rows.Scan(&interaction); err != nil {
			return nil, err
		}

		interactions = append(interactions, interaction)
	}

	// Return the created user
	return interactions, err
}

