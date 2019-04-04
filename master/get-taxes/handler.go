package gettaxes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	utilctx "github.com/naufalridho/tax-calculator/common/context"
)

func GetTaxesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := utilctx.NewContextFromRequest(r)

	var err error
	form := NewGetTaxesForm()

	form.Response = GetTaxesResponse{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
	}

	defer func() {
		if err != nil {
			form.Response.Message = err.Error()
			log.Println(err)
		}
		data, _ := json.Marshal(form.Response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}()

	err = form.ValidateRequest(ctx, r)
	if err != nil {
		return
	}

	err = form.GetTaxesService(ctx)
	if err != nil {
		return
	}

	form.Response = GetTaxesResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       form.BillList,
	}
	return
}
