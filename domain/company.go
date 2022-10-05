package domain

import (
	"database/sql"
	"time"
	"github.com/rs/zerolog/log"
)

type Company struct {
	Id string
	Name string
	Created time.Time
}

type CompanyModel struct {
	DB *sql.DB
}

func (m CompanyModel) All() ([]Company,error) {
	rows, err := m.DB.Query("select id,name,created from Companys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companys []Company

	for rows.Next() {
		var c Company

		err := rows.Scan(&c.Id, &c.Name, &c.Created)
		if err != nil {
			return nil, err
		}
		// log.Info().Msg(c.id)
		companys = append(companys, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companys, nil
}
