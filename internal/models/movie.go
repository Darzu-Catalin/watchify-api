package models

import (
	"time"
)
type CustomDate struct {
    time.Time
}

// UnmarshalJSON customizes the JSON unmarshaling for CustomDate
func (c *CustomDate) UnmarshalJSON(data []byte) error {
    str := string(data)
    t, err := time.Parse(`"2006-01-02"`, str)
    if err != nil {
        return err
    }
    c.Time = t
    return nil
}
type Movie struct {
    ID          int       `json:"id,omitempty"`
    Title       string    `json:"title"`
    ReleaseDate string `json:"release_date"`
    Overview    string    `json:"overview"`
    Rating      float64   `json:"rating"`
    Popularity  float64   `json:"popularity"`
    Runtime     int       `json:"runtime"`
    Budget      int64     `json:"budget"`
    Revenue     int64     `json:"revenue"`
    PosterPath  string    `json:"poster_path"`
}