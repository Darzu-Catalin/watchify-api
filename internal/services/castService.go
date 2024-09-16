package services

import (
    "WatchifyAPI/internal/db"
    "WatchifyAPI/internal/models"
    
)

func GetCastByMovieId(movieId int) ([]models.Cast, error) {
	var castMembers []models.Cast

    // Use Query to retrieve multiple rows
    rows, err := db.DB.Query("SELECT * FROM cast WHERE movie_id = ?", movieId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var cast models.Cast
        err := rows.Scan(&cast.ID, &cast.Name, &cast.CharacterName, &cast.Movie_id)
        if err != nil {
            return nil, err
        }
        castMembers = append(castMembers, cast)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return castMembers, nil
}