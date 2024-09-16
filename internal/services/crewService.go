package services

import (
    "WatchifyAPI/internal/db"
    "WatchifyAPI/internal/models"
    
)

func GetCrewByMovieId(movieId int) ([]models.Crew, error) {
	var crewMembers []models.Crew

    // Use Query to retrieve multiple rows
    rows, err := db.DB.Query("SELECT * FROM crew WHERE movie_id = ?", movieId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var crew models.Crew
        err := rows.Scan(&crew.ID, &crew.Name, &crew.Job, &crew.Movie_id)
        if err != nil {
            return nil, err
        }
        crewMembers = append(crewMembers, crew)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return crewMembers, nil
}