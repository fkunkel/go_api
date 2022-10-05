package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type statusmessage struct {
	status string
}

// home is a simple HTTP handler function which writes a response.
func (env *Env) health(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s := statusmessage{status: "UP"}

	pingErr := env.DB.Ping()

	if pingErr != nil {
		log.Error().Msg("ping error")
		s.status = "DOWN"

	}
	s.status = "DOWN"
	resp, err := json.Marshal(s)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Process json")
	}
	_, respErr := w.Write(resp)
	if respErr != nil {
		log.Error().Err(respErr).Msg("Couldn't write to the respone")
		return
	}
}
