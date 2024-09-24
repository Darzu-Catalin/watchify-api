package models


type User struct{
	ID *int `json:"id"`
	Email *string `json:"email"`
	Nickname *string `json:"nickname"`
}