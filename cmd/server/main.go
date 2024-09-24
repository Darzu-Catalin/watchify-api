package main

import (
	"WatchifyAPI/internal/db"
	"WatchifyAPI/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

// CORS middleware function
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, you can change this to specific domains
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize the database
	if err := db.InitDB("root:@tcp(127.0.0.1:3306)/watchify"); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register your handlers
	mux.HandleFunc("/movie", handlers.GetMovieById)
	mux.HandleFunc("/movie/add", handlers.AddMovie)
	mux.HandleFunc("/movie/delete", handlers.DeleteMovie)
	mux.HandleFunc("/crew", handlers.GetCrewByMovieId)
	mux.HandleFunc("/cast", handlers.GetCastByMovieId)
	mux.HandleFunc("/genre", handlers.GetGenreByMovieId)
	mux.HandleFunc("/video", handlers.GetVideoByMovieId)
	mux.HandleFunc("/production", handlers.GetProductionByMovieId)
	mux.HandleFunc("/getData", handlers.GetMovieDataByMovieId)
	mux.HandleFunc("/getAllGenres", handlers.GetAllGenres)
	mux.HandleFunc("/searchByName", handlers.SearchMovieByName)
	mux.HandleFunc("/searchByActor", handlers.SearchMovieByActor)
	mux.HandleFunc("/getNewMovies", handlers.GetNewMoviesByUserId)
	mux.HandleFunc("/getTrendingMovies", handlers.GetTrandingMoviesByUserId)
	mux.HandleFunc("/getMoviesByGenre", handlers.GetMoviesByGenre)
	mux.HandleFunc("/addUser", handlers.AddUser)
	mux.HandleFunc("/addUserInteraction", handlers.AddUserInteraction)

	// Wrap the mux with the CORS middleware
	corsHandler := enableCORS(mux)

	// Start the server
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", corsHandler))
	fmt.Println("Success!")
}
