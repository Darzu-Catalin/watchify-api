package handlers

import (
	"WatchifyAPI/internal/models"
	"WatchifyAPI/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	// "WatchifyAPI/internal/models"
	// "log"
	// "WatchifyAPI/internal/db"
)

func GetMovieDataByMovieId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	movieId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	movieData, err := services.GetMovieDataByMovieId(movieId)
	if err != nil {
		http.Error(w, "Failed to retrieve crew members", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetNewMoviesByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}
	limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}

	movieData, err := services.GetNewMoviesByUserId(userId, start, stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetTrandingMoviesByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}

	movieData, err := services.GetTrandingMoviesByUserId(userId, start, stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	genres, ok := r.URL.Query()["genre"]
	if !ok || len(genres[0]) < 1 {
		http.Error(w, "Missing 'genre' parameter", http.StatusBadRequest)
		return
	}

    limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}


	movieData, err := services.GetMoviesByGenre(genres[0],start,stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetWatchLaterMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

    limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}


	movieData, err := services.GetWatchLaterMovies(userId,start,stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetLikedMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

    limitStart, ok := r.URL.Query()["start"]
	if !ok || len(limitStart[0]) < 1 {
		http.Error(w, "Missing 'start' parameter", http.StatusBadRequest)
		return
	}

	start, err := strconv.Atoi(limitStart[0])
	if err != nil {
		http.Error(w, "Invalid 'start' parameter", http.StatusBadRequest)
		return
	}

	limitStop, ok := r.URL.Query()["stop"]
	if !ok || len(limitStop[0]) < 1 {
		http.Error(w, "Missing 'stop' parameter", http.StatusBadRequest)
		return
	}

	stop, err := strconv.Atoi(limitStop[0])
	if err != nil {
		http.Error(w, "Invalid 'stop' parameter", http.StatusBadRequest)
		return
	}


	movieData, err := services.GetLikedMovies(userId,start,stop)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// func GetLikedMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	ids, ok := r.URL.Query()["id"]
// 	if !ok || len(ids[0]) < 1 {
// 		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
// 		return
// 	}

// 	userId, err := strconv.Atoi(ids[0])
// 	if err != nil {
// 		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
// 		return
// 	}

// 	moviesId, err := services.GetLikedMovies(userId)
// 	if err != nil {
// 		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
// 		return
// 	}
// 	var movieData []models.Universal
	


// 	jsonData, err := json.Marshal(movieData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write(jsonData)
// }


type MovieRecommendation struct {
	Results []struct {
		ID int `json:"id"`
	} `json:"results"`
}

func GetRecMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieData []*models.Universal

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}



	id, err := services.GetLikedMoviesId(userId)
	if err != nil{
		http.Error(w, "Error getting liked movies"+err.Error(), http.StatusBadRequest)
		return
	}

	str := strconv.Itoa(id)

	url := "https://api.themoviedb.org/3/movie/" + str + "/recommendations?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiZWFmNzAxYWNjMDU4ZTlmZGNjOTc0ODI4MjYyMWU3NyIsIm5iZiI6MTcyNzM0MTMyMy45OTIwOTYsInN1YiI6IjY2ZGIwNTdmNjg2M2M0Zjg5ZTM0ZDc0MiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.A2-egCDgsA3hR621YNG7fflstmH3cjti6oMlXOyunEs")

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// Create a struct to hold the parsed data
	var recommendations MovieRecommendation
	err = json.Unmarshal(body, &recommendations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}



	// Extract and print only the IDs from the results
	for _, movie := range recommendations.Results {
		var temp *models.Universal
		temp, err = services.GetMovieDataByMovieId(movie.ID)
		if err != nil{
			log.Fatal("Error while getting data")
		}
		movieData = append(movieData, temp)
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}


func GetSimilarMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieData []*models.Universal

	ids, ok := r.URL.Query()["movieId"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}


	

	url := "https://api.themoviedb.org/3/movie/" + ids[0] + "/similar?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiZWFmNzAxYWNjMDU4ZTlmZGNjOTc0ODI4MjYyMWU3NyIsIm5iZiI6MTcyNzM0MTMyMy45OTIwOTYsInN1YiI6IjY2ZGIwNTdmNjg2M2M0Zjg5ZTM0ZDc0MiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.A2-egCDgsA3hR621YNG7fflstmH3cjti6oMlXOyunEs")

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// Create a struct to hold the parsed data
	var recommendations MovieRecommendation
	err = json.Unmarshal(body, &recommendations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}



	// Extract and print only the IDs from the results
	for _, movie := range recommendations.Results {
		var temp *models.Universal
		temp, err = services.GetMovieDataByMovieId(movie.ID)
		if err != nil{
			log.Fatal("Error while getting data")
		}
		movieData = append(movieData, temp)
	}

	jsonData, err := json.Marshal(movieData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

