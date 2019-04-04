package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (d *DB) Connect() error {
	db, err := sqlx.Connect(d.Type, d.Conn)
	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("[Error]: DB ping connection error", err.Error())
		return err
	}

	d.DBConnection = db

	return nil
}
