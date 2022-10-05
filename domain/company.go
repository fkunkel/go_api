package domain

import (
	"database/sql"
	"time"
	"github.com/rs/zerolog/log"
)

type Company struct {
	id string
	name string
	created time.Duration
}

type CompanyModel struct {
	DB *sql.DB
}

func (m CompanyModel) All() ([]Company,error) {
	rows, err := m.DB.Query("SELECT * FROM Companys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companys []Company

	for rows.Next() {
		var c Company

		err := rows.Scan(&c.id, &c.name, &c.created)
		if err != nil {
			return nil, err
		}

		companys = append(companys, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	log.Info().Msg(string(len(companys)))
	return companys, nil
}
