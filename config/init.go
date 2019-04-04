package config

import (
	_ "github.com/lib/pq"
	"gopkg.in/gcfg.v1"
	"log"
)

var Cfg struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host string
}

func init() {
	path := "etc/tax-calculator/tax-calculator"
	fname := "taxcalculator.development.ini"
	err := gcfg.ReadFileInto(&Cfg, path+"/"+fname)
	if err != nil {
		log.Fatalln("Cannot read config. Err:", err)
	}
	Cfg.Server.Host = ":8080"
}
