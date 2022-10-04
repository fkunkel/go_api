package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)
type statusmessage struct {
	Status string
}
// home is a simple HTTP handler function which writes a response.
func(env *Env) health(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err :=json.Marshal(statusmessage{"UP"})
	if err != nil {
		log.Error().Err(err).Msg("Unable to Process json")
	}
	w.Write(resp)
	}
