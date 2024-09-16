package models


type MovieGenres struct {
	Movie_id int `json:"movie_id"`
	Genre 	 []Genres `json:"genre"`
	Genre_id int `json:"genre_id"`
}