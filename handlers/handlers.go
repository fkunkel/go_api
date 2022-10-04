package handlers

import (
	"github.com/fkunkel/go_api/domain"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Env struct {
	logger   *zerolog.Logger
	companys interface {
		All() ([]domain.Company, error)
	}
}

func (env *Env) Routers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/company", env.companyAll).Methods("GET")

	return r
}
