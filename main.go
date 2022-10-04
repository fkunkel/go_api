package main

import (
	"database/sql"
	"time"

	// "github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/fkunkel/go_api/handlers"
	"github.com/fkunkel/go_api/domain"
	"net/http"
)

type Env struct {
	db *sql.DB
	logger *zerolog.Logger
}
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.Logger{}

	db, err := sql.Open("mysql", "platformUser:Wombat2016#@/platfo")
	if err != nil {
		logger.Error().Msg("Cannot make connection")
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := handlers.Routers()
	log.Err(http.ListenAndServe(":8000", router))

}
