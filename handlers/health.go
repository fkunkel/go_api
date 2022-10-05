package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type statusMessage struct {
	Status string
}

// home is a simple HTTP handler function which writes a response.
func (env *Env) health(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")

	statusUp := "UP"
	httpStatus := http.StatusOK

	pingErr := env.DB.Ping()

	if pingErr != nil {
		log.Error().Msg("ping error")
		statusUp = "DOWN"
		httpStatus = http.StatusServiceUnavailable
	}

	resp, err := json.Marshal(statusMessage{Status: statusUp})
	w.WriteHeader(httpStatus)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Process json")
	}
	_, respErr := w.Write(resp)
	if respErr != nil {
		log.Error().Err(respErr).Msg("Couldn't write to the response")
		return
	}
}
