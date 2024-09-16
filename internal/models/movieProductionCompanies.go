package models


type MovieProductionCompanies struct{
	Movie_id int `json:"movie_id"`
	Name []ProductionCompanies `json:"production_companies_name"`
	Company_id int `json:"company_id"`
}