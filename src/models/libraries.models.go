package models

import "time"

type Libraries struct {
	Id_libraries int `json:"id_libraries,omitempty"`
	Id_users int `json:"id_users,omitempty"`
	Name string `json:"name,omitempty"`
	Path int `json:"path,omiempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Scanned_at time.Time `json:"scanned_at,omitempty"`
}