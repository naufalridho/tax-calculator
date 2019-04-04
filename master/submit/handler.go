package submit

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	utilctx "github.com/naufalridho/tax-calculator/common/context"
	"github.com/naufalridho/tax-calculator/common/format"
)

func SubmitTaxHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := utilctx.NewContextFromRequest(r)

	var err error
	form := NewSubmitForm()
	form.Response = SubmitResponse{
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
		w.WriteHeader(form.Response.StatusCode)
		w.Write(data)
	}()

	form.Request = SubmitRequest{
		Name:  r.FormValue("name"),
		Code:  format.ToInt(r.FormValue("code")),
		Price: format.ToFloat64(r.FormValue("price")),
	}

	err = form.ValidateRequest(ctx, r)
	if err != nil {
		return
	}

	err = form.SubmitService(ctx)
	if err != nil {
		return
	}

	form.Response = SubmitResponse{
		Success:    true,
		StatusCode: http.StatusOK,
	}
	return
}
