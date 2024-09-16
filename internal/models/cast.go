package models

type Cast struct{
	ID   		   int 	   `json:"id"`
	Name 		   string  `json:"name"`
	CharacterName  string  `json:"character_name"`
	Movie_id	   int 	   `json:"movie_id"`
}