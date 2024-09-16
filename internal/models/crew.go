package models

type Crew struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
	Movie_id int `json:"movie_id"`
}