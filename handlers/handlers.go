package handlers

import (
	"database/sql"
	"github.com/fkunkel/go_api/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

type Env struct {
	DB *sql.DB //Use this to get direct access if needed
	companys interface {
		All() ([]domain.Company, error)
	}
}

type Config struct {
	DBDriver      string `json:"DB_DRIVER"`
	DBSource      string `json:"DB_SOURCE"`
	ServerAddress string `json:"SERVER_ADDRESS"`
}

func(c *Config) ConfigService() (*Env,error) {

	db, dbErr := c.dbService()
	if dbErr != nil {  return nil, dbErr }

	env := &Env{DB:db, companys: domain.CompanyModel{DB: db}}
	//defer db.Close()

	return env, nil
}
func (env *Env) Routers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", env.health).Methods("GET")
	r.HandleFunc("/company", env.companyAll).Methods("GET")

	return r
}

func(c *Config) dbService() (*sql.DB, error){
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Made it")
	db, err := sql.Open(c.DBDriver,c.DBSource )
	if err != nil {
		log.Error().Msg("Cannot make connection")
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
