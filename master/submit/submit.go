package submit

import (
	"context"
	"fmt"
	"net/http"

	bills "github.com/naufalridho/tax-calculator/class/bill"
	utilctx "github.com/naufalridho/tax-calculator/common/context"
	"github.com/naufalridho/tax-calculator/common/database"
	"github.com/naufalridho/tax-calculator/state"
)

type SubmitForm struct {
	Request  SubmitRequest
	Response SubmitResponse
}

type SubmitRequest struct {
	Name  string  `db:"name" json:"name"`
	Code  int     `db:"tax_code" json:"code"`
	Price float64 `db:"price" json:"price"`
}

type SubmitResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewSubmitForm() *SubmitForm {
	return &SubmitForm{}
}

func (form *SubmitForm) ValidateRequest(ctx context.Context, r *http.Request) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	if r.Method != http.MethodPost {
		form.Response.StatusCode = http.StatusMethodNotAllowed
		return fmt.Errorf("Invalid method")
	}

	if form.Request.Name == "" {
		form.Response.StatusCode = http.StatusBadRequest
		return fmt.Errorf("Name field cannot be empty")
	}
	if _, ok := state.TaxCodeName[form.Request.Code]; !ok {
		form.Response.StatusCode = http.StatusBadRequest
		return fmt.Errorf("Tax code is invalid")
	}

	return nil
}

func (form *SubmitForm) SubmitService(ctx context.Context) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	var err error

	err = form.InsertBill(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (form *SubmitForm) InsertBill(ctx context.Context) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	err := bills.InsertBill(ctx, database.PostgresDB.DBConnection,
		form.Request.Name,
		form.Request.Code,
		form.Request.Price,
	)
	if err != nil {
		return err
	}

	return nil
}
