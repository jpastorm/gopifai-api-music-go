package apihelpers

import "wopifai/src/resources"

type ResponseLibraries struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []resources.Explorer `json:"data"`
}
