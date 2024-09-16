package models


type Universal struct {
	Movie *Movie      `json:"movie"`
	Genre *MovieGenres `json:"genres"`
	Production *MovieProductionCompanies `json:"production"`
	Video []Video    `json:"video_path"`
	Cast  []Cast `json:"crew_cast"`
	Crew  []Crew   `json:"crew"`

}