package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqlProperty struct {
	Conn *sqlx.DB
}

func NewSqlProperty(driverName string, dataSourceName string) (*SqlProperty, error) {
	var p SqlProperty
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p.Conn = db

	return &p, nil
}
