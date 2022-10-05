package handlers

import (
	"database/sql"

	"github.com/fkunkel/go_api/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"time"
)

type Env struct {
	logger   *zerolog.Logger
	companys interface {
		All() ([]domain.Company, error)
	}
}

func ConfigService() *Env {
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
	pingErr := db.Ping()

	if pingErr != nil {
		logger.Error().Msg("ping error")
	}

	env := &Env{logger: &logger,companys: domain.CompanyModel{db}}
	//defer db.Close()

	return env
}
func (env *Env) Routers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/company", env.companyAll).Methods("GET")

	return r
}

