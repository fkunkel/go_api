package main

import (
	"encoding/json"
	"github.com/fkunkel/go_api/handlers"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	c := handlers.Config{}
	file, err := os.Open("app.json")
	if err != nil {
		log.Error().Err(err).Msg("Cannot find file")
	return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Error().Err(err).Msg("Cannot decode file")
		return
	}
	env,err := c.ConfigService()
	if err != nil {
		log.Error().Err(err).Msg("error in config")
	}
	router := env.Routers()
	log.Err(http.ListenAndServe(c.ServerAddress, router))

}
