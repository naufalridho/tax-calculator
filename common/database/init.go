package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/gcfg.v1"
)

//Parameters got from config
var DBConfig struct {
	PostgresCfg Config
}

type Config struct {
	Conn string
	Type string
}

//Primary Database ObjectF
type DB struct {
	DBConnection *sqlx.DB
	Conn         string
	Type         string
}

var PostgresDB *DB

func init() {
	path := "etc/tax-calculator/database"
	fname := "database.development.ini"
	err := gcfg.ReadFileInto(&DBConfig, path+"/"+fname)
	if err != nil {
		log.Fatalln("Cannot read db config. Err:", err)
	}

	PostgresDB = &DB{
		Conn: DBConfig.PostgresCfg.Conn,
		Type: DBConfig.PostgresCfg.Type,
	}
	err = PostgresDB.Connect()
	if err != nil {
		log.Fatalln("Failed to connect to db. Err:", err)
	}
}
