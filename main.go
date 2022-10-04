package main

import (
	"database/sql"
	"github.com/fkunkel/go_api/domain"
	"github.com/fkunkel/go_api/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.Logger{}
	logger.Info().Msg("Made it here")
	db, err := sql.Open("mysql", "platformUser:Wombat2016#@/platformtest")
	if err != nil {
		logger.Error().Msg("Cannot make connection")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	env := &handlers.Env{logger: &logger,companys: domain.CompanyModel{db}}

	router := env.handlers.Routers()
	log.Err(http.ListenAndServe(":8000", router))

}
