package services

import (
	// "WatchifyAPI/internal/db"
	"WatchifyAPI/internal/models"
	// "net/http"
	"WatchifyAPI/internal/db"
)

func SearchByName(prompt string, start int, stop int) ([]*models.Universal, error) {
	var foundMovies []*models.Universal
	var ids []int

	// Query to get movie IDs
	query := "SELECT id FROM movies WHERE title LIKE ? LIMIT ?,?"
	searchTerm := "%" + prompt + "%"
	rows, err := db.DB.Query(query, searchTerm, start, stop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	for _, value := range ids {
		movie, err := GetMovieDataByMovieId(value)
		if err != nil {
			return nil, err
		}
		foundMovies = append(foundMovies, movie)
	}

	return foundMovies, nil
}

func SearchByActor(name string, start int, stop int) ([]*models.Universal, error) {
	var foundMovies []*models.Universal
	var ids []int
	query := "SELECT movie_id FROM cast WHERE name LIKE ? LIMIT ?,?"
	searchTerm := "%" + name + "%"
	rows, err := db.DB.Query(query, searchTerm, start, stop)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	for _, value := range ids {
		movie, err := GetMovieDataByMovieId(value)
		if err != nil {
			return nil, err
		}
		foundMovies = append(foundMovies, movie)
	}

	return foundMovies, nil
}
