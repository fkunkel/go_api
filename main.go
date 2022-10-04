package main

import (
	"github.com/fkunkel/go_api/handlers"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	router := handlers.Routers()
	log.Err(http.ListenAndServe(":8000", router))

}
