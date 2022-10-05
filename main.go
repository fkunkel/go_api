package main

import (
	"github.com/fkunkel/go_api/handlers"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	env,err := handlers.ConfigService()
	if err != nil {
		log.Error().Err(err).Msg("error in config")
	}
	router := env.Routers()
	log.Err(http.ListenAndServe(":8000", router))

}
