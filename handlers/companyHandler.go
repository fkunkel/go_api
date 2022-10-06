package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

// company is a simple HTTP handler function which writes a response.
func (env *Env) companyAll(w http.ResponseWriter, _ *http.Request) {
	companys,err := env.companys.All()

	if err != nil {
		log.Error().Err(err).Msg("Couldn't get companies")
	}

	resp, jsonErr :=json.Marshal(companys)
	if jsonErr != nil {
		log.Error().Err(jsonErr).Msg("Unable to Process json")
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Error().Err(err).Msg("write error")
	}

}
