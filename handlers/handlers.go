package handlers

import (
	"database/sql"
	"github.com/rs/zerolog/log"

	"github.com/fkunkel/go_api/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"time"
)

type Env struct {
	DB *sql.DB //Use this to get direct access if needed
	companys interface {
		All() ([]domain.Company, error)
	}
}

func ConfigService() (*Env,error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Made it")
	db, err := sql.Open("mysql", "platformUser:Wombat2016#@/platformtest?parseTime=true")
	if err != nil {
		log.Error().Msg("Cannot make connection")
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)


	env := &Env{DB:db, companys: domain.CompanyModel{db}}
	defer db.Close()

	return env, nil
}
func (env *Env) Routers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", env.health).Methods("GET")
	r.HandleFunc("/company", env.companyAll).Methods("GET")

	return r
}
