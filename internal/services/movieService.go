package services

import (
    "WatchifyAPI/internal/db"
    "WatchifyAPI/internal/models"
    
)

func GetMovieById(id int) (*models.Movie, error) {
    var movie models.Movie
    var releaseDateStr string

    err := db.DB.QueryRow("SELECT * FROM movies WHERE id = ? and poster_path is not null", id).
        Scan(
			&movie.ID, 
			&movie.Title, 
			&releaseDateStr, 
			&movie.Overview, 
			&movie.Rating, 
			&movie.Popularity, 
			&movie.Runtime, 
			&movie.Budget, 
			&movie.Revenue, 
			&movie.PosterPath,
		)

    if err != nil {
        return nil, err
    }

    // movie.ReleaseDate, _ = time.Parse("2006-01-02", releaseDateStr)
	movie.ReleaseDate = releaseDateStr
    return &movie, nil
}

func AddMovie(movie models.Movie) error {
    // Prepare the SQL statement
    stmt, err := db.DB.Prepare(`
        INSERT INTO movies (title, release_date, overview, rating, popularity, runtime, budget, revenue, poster_path)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the statement with movie data
    _, err = stmt.Exec(
        movie.Title,
        movie.ReleaseDate,
		// Format("2006-01-02"), // Ensure the format matches your DB schema
        movie.Overview,
        movie.Rating,
        movie.Popularity,
        movie.Runtime,
        movie.Budget,
        movie.Revenue,
        movie.PosterPath,
    )
    return err
}





