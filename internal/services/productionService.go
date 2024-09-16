package services

import (
    "WatchifyAPI/internal/db"
    "WatchifyAPI/internal/models"
    
)

func GetProductionByMovieId(movieId int) (*models.MovieProductionCompanies, error) {
	var production models.MovieProductionCompanies
	// var genres []models.Genres
	var err error
    // Use Query to retrieve multiple rows
    rows, err := db.DB.Query("SELECT id, name FROM production_companies pc inner join movie_production_companies mpc on pc.id = mpc.company_id WHERE movie_id = ?", movieId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var productionCompany models.ProductionCompanies
        err := rows.Scan(&productionCompany.ID, &productionCompany.Name)
        if err != nil {
            return nil, err
        }
        production.Name = append(production.Name, productionCompany)
    }

	if err = rows.Err(); err != nil {
        return nil, err
    }

	erro := db.DB.QueryRow("SELECT movie_id, company_id FROM movie_production_companies mpc inner join production_companies pc on mpc.company_id = pc.id WHERE movie_id = ? LIMIT 1", movieId).
        Scan(
			&production.Movie_id, 
			&production.Company_id,
					
		)
	if erro != nil {
		return nil, err
	}
	
	

    // Check for errors from iterating over rows
    

    return &production, nil
}