package api

import (
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	mh := new(MultiplyHandler)
	r.HandleFunc("/v1/multiply", mh.Post).Methods("POST")
}