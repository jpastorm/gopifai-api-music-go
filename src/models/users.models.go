package models

import "time"

type Users struct {
	Id_users int    `json:"id_users,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email string `json:"email,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
	State bool `json:"state,omitempty"`
	Last_login time.Time `json:"state,omitempty"`
}
