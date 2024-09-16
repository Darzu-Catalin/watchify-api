package main

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/handlers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "encoding/json"
	// "strconv"
	// "github.com/gin-gonic/gin"
	// "errors"
	// "time"
	// "database/sql"
)

// func getMovieById(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	ids, ok := r.URL.Query()["id"]
//     if !ok || len(ids[0]) < 1 {
//         http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
//         return
//     }
// 	id, err := strconv.Atoi(ids[0])
//     if err != nil {
//         http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
//         return
//     }

// 	var movie Movie
// 	var releaseDateStr string

// 	query := "Select * from movies where id = ?"

// 	err = db.QueryRow(query, id).Scan(
// 		&movie.ID,
//         &movie.Title,
//         &releaseDateStr,
//         &movie.Overview,
//         &movie.Rating,
//         &movie.Popularity,
//         &movie.Runtime,
//         &movie.Budget,
//         &movie.Revenue,
//         &movie.PosterPath,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
//             http.Error(w, "Movie not found", http.StatusNotFound)
//         } else {
//             log.Printf("Error retrieving movie: %v", err)
//             http.Error(w, "Server error", http.StatusInternalServerError)
//         }
//         return
// 	}

// 	movie.ReleaseDate, err = time.Parse("2006-01-02", releaseDateStr)
// 	if err != nil {
// 		log.Printf("Error parsing date: %v", err)
//         http.Error(w, "Invalid date format", http.StatusInternalServerError)
//         return
// 	}

// 	jsonData, err := json.Marshal(movie)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     // Write the JSON response
//     w.Write(jsonData)
// }




func main() {
	if err := db.InitDB("root:@tcp(127.0.0.1:3306)/watchify"); err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
	http.HandleFunc("/movie", handlers.GetMovieById)
    http.HandleFunc("/movie/add", handlers.AddMovie)
    http.HandleFunc("/movie/delete", handlers.DeleteMovie)
    http.HandleFunc("/crew", handlers.GetCrewByMovieId)
    http.HandleFunc("/cast", handlers.GetCastByMovieId)
    http.HandleFunc("/genre", handlers.GetGenreByMovieId)
	http.HandleFunc("/video", handlers.GetVideoByMovieId)
	http.HandleFunc("/production", handlers.GetProductionByMovieId)
	http.HandleFunc("/getData", handlers.GetMovieDataByMovieId)


    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
    fmt.Println("Success!")

}


