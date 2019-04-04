package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/naufalridho/tax-calculator/config"
	"github.com/naufalridho/tax-calculator/router"
)

func main() {
	r := httprouter.New()
	router.InitTaxCalculatorRouter(r)

	log.Println("service is running...")
	log.Fatal(http.ListenAndServe(config.Cfg.Server.Host, r))
}
