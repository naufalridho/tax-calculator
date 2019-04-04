package router

import (
	"github.com/julienschmidt/httprouter"
	deletetax "github.com/naufalridho/tax-calculator/master/delete-tax"
	gettaxes "github.com/naufalridho/tax-calculator/master/get-taxes"
	"github.com/naufalridho/tax-calculator/master/submit"
)

func InitTaxCalculatorRouter(r *httprouter.Router) {
	r.POST("/v1/taxes", submit.SubmitTaxHandler)
	r.GET("/v1/taxes", gettaxes.GetTaxesHandler)
	r.DELETE("/v1/taxes", deletetax.DeleteTaxHandler)
}
