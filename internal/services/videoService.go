package services

import (
    "WatchifyAPI/internal/db"
    "WatchifyAPI/internal/models"
    
)

func GetVideoByMovieId(movieId int) ([]models.Video, error) {
	var videos []models.Video

    // Use Query to retrieve multiple rows
    rows, err := db.DB.Query("SELECT * FROM videos WHERE movie_id = ?", movieId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var video models.Video
        err := rows.Scan(&video.ID, &video.Movie_id, &video.Name, &video.VType, &video.Url)
        if err != nil {
            return nil, err
        }
        videos = append(videos, video)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return videos, nil
}