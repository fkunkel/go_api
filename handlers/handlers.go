package handlers

import "github.com/gorilla/mux"


func  Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", health).Methods("GET")
	return r
}
