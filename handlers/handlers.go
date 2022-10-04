package handlers

import (
	"database/sql"
	"github.com/fkunkel/go_api/domain"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"time"
)
type Env struct {
	logger *zerolog.Logger
	companys interface{
		All() ([]domain.Company, error)
	}
}

func Routers() *mux.Router {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.Logger{}

	db, err := sql.Open("mysql", "platformUser:Wombat2016#@/platformtest")
	if err != nil {
		logger.Error().Msg("Cannot make connection")
	}
	env := &Env{logger: &logger,companys: domain.CompanyModel{db}}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/company", env.companyAll).Methods("GET")

	return r
}
