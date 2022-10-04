package main

import (
	 // "github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/fkunkel/go_api/handlers"
	"net/http"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Starting the service...")

	router := handlers.Routers()
	log.Err(http.ListenAndServe(":8000", router))

}
