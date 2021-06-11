package main

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func router() *mux.Router {
	router := mux.NewRouter()

	//app info route
	router.HandleFunc("/api/heartbeat", getInfo).Methods("GET", "OPTIONS")
}
