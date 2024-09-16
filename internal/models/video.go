package models


type Video struct{
	ID 		 string   `json:"id"`
	Movie_id int 	  `json:"movie_id"`
	Name     string   `json:"name"` 
	VType	 string   `json:"type"`
	Url 	 string	  `json:"url"`
}