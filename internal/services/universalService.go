package services

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/models"

	// "net/http"
	"log"
)

func GetMovieDataByMovieId(movieId int) (*models.Universal, error) {
	var movieData models.Universal
	var err error

	movieData.Movie, err = GetMovieById(movieId)
	if err != nil {
		// handle the error

		log.Println(err) // or return the error
		
	}

	movieData.Genre, err = GetGenreByMovieId(movieId)
	if err != nil {
		// handle the error
		log.Println(err)
		 // or return the error
	}

	movieData.Production, err = GetProductionByMovieId(movieId)
	if err != nil {
		// handle the error
		log.Println(err)
		 // or return the error
	}

	movieData.Video, err = GetVideoByMovieId(movieId)
	if err != nil {
		// handle the error
		log.Println(err) // or return the error
	}

	movieData.Cast, err = GetCastByMovieId(movieId)
	if err != nil {
		log.Println(err)
	}

	movieData.Crew, err = GetCrewByMovieId(movieId)
	if err != nil {
		log.Println(err)
	}
	// if err != nil {
	//     http.Error(w, "Failed to retrieve crew members", http.StatusInternalServerError)
	//     return
	// }

	// var genres []models.Genres
	// var err error
	// // Use Query to retrieve multiple rows
	// rows, err := db.DB.Query("SELECT id, name FROM production_companies pc inner join movie_production_companies mpc on pc.id = mpc.company_id WHERE movie_id = ?", movieId)
	// if err != nil {
	//     return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	//     var productionCompany models.ProductionCompanies
	//     err := rows.Scan(&productionCompany.ID, &productionCompany.Name)
	//     if err != nil {
	//         return nil, err
	//     }
	//     production.Name = append(production.Name, productionCompany)
	// }

	// if err = rows.Err(); err != nil {
	//     return nil, err
	// }

	// erro := db.DB.QueryRow("SELECT movie_id, company_id FROM movie_production_companies mpc inner join production_companies pc on mpc.company_id = pc.id WHERE movie_id = ? LIMIT 1", movieId).
	//     Scan(
	// 		&production.Movie_id,
	// 		&production.Company_id,

	// 	)
	// if erro != nil {
	// 	return nil, err
	// }

	// Check for errors from iterating over rows

	return &movieData, nil
}

func GetNewMoviesByUserId(id int, limitStart int, limitFinish int) ([]*models.Universal, error) {
	var movies []*models.Universal
	var ids []int
	query := "SELECT id FROM movies m WHERE NOT EXISTS (SELECT 1 FROM user_movie_interactions ui WHERE ui.movie_id = m.id AND ui.user_id = ? and m.poster_path is not null) AND m.release_date < CURRENT_DATE ORDER BY m.release_date DESC LIMIT ?, ?"

	rows, err := db.DB.Query(query, id, limitStart, limitFinish)
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
		movies = append(movies, movie)
	}

	return movies, nil
}

func GetTrandingMoviesByUserId(id int, limitStart int, limitFinish int) ([]*models.Universal, error) {
	var movies []*models.Universal
	var ids []int
	query := "SELECT id FROM movies m WHERE NOT EXISTS (SELECT 1 FROM user_movie_interactions ui WHERE ui.movie_id = m.id AND ui.user_id = ?) AND m.release_date < CURRENT_DATE ORDER BY popularity DESC,m.release_date DESC LIMIT ?,?"

	rows, err := db.DB.Query(query, id,limitStart,limitFinish)
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
		movies = append(movies, movie)
	}

	return movies, nil
}


func GetMoviesByGenre(genre string, limitStart int, limitFinish int) ([]*models.Universal, error) {
	var movies []*models.Universal
	var ids []int
	query := "SELECT movie_id FROM movie_genres m inner join genres g on m.genre_id = g.id WHERE g.name = ? LIMIT ?,? "

	rows, err := db.DB.Query(query, genre, limitStart, limitFinish)
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
		movies = append(movies, movie)
	}

	return movies, nil
}
