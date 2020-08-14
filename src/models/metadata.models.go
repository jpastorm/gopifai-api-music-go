package models

import "time"

type Metadata struct {
	Id_metadata int `json:"id_metadata,omitempty"`
	Id_multimedia int `json:"id_multimedia,omitempty"`
	Artist string `json:"artist,omitempty"`
	Album string `json:"album,omitempty"`
	Track string `json:"track,omitempty"`
	Genre string `json:"genre,omitempty"`
	Year string `json:"year,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}
