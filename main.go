package main

import (
	"fmt"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/fkunkel/go_api/handlers"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Starting the service...")

	router := handlers.Router()

}
