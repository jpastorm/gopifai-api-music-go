package models

import "time"

type Multimedia struct {
	Id_multimedia string `json:"id_multimedia,omitempty"`
	Id_libraries string `json:"id_libraries,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Created_at time.Time `json:"created_at,omiempty"`
	Lastview_at time.Time `json:"lastview_at,omiempty"`
	Likeit bool `json:"likeit,omiempty"`
}
