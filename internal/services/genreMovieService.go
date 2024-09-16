package services

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/models"
    "fmt"

	
)

func GetGenreByMovieId(movieId int) (*models.MovieGenres, error) {
	var movieGenre models.MovieGenres
	// var genres []models.Genres
	var err error
    // Use Query to retrieve multiple rows
    rows, err := db.DB.Query("SELECT id, name FROM movie_genres inner join genres on movie_genres.genre_id = genres.id WHERE movie_id = ?", movieId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var genre models.Genres
        err := rows.Scan(&genre.ID, &genre.Name)
        if err != nil {
            return nil, err
        }
        movieGenre.Genre = append(movieGenre.Genre, genre)
    }

	if err = rows.Err(); err != nil {
        return nil, err
    }

	erro := db.DB.QueryRow("SELECT movie_id, genre_id FROM movie_genres inner join genres on movie_genres.genre_id = genres.id WHERE movie_id = ? LIMIT 1", movieId).
        Scan(
			&movieGenre.Movie_id, 
			&movieGenre.Genre_id,
					
		)
	if erro != nil {
		return nil, err
	}
	
	

    // Check for errors from iterating over rows
    

    return &movieGenre, nil
}

func GetAllGenres() ([]string, error){
    var genres []string
    var err error

    rows, err := db.DB.Query("SELECT name FROM genres")
    
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next(){
        var genreName string;
        err := rows.Scan(&genreName)
        if err != nil {
            return nil, err
        }
        genres = append(genres, genreName)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    fmt.Println("Array:", genres)

    return genres, nil
}