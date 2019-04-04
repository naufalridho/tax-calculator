package gettaxes

import (
	"context"
	"fmt"
	bills "github.com/naufalridho/tax-calculator/class/bill"
	utilctx "github.com/naufalridho/tax-calculator/common/context"
	"github.com/naufalridho/tax-calculator/common/database"
	"github.com/naufalridho/tax-calculator/state"
	"net/http"
)

type GetTaxesForm struct {
	BillList      []*bills.Bill
	PriceSubtotal float64
	TaxSubtotal   float64
	GrandTotal    float64

	Response GetTaxesResponse
}

type GetTaxesResponse struct {
	Success    bool          `json:"success"`
	Message    string        `json:"message"`
	StatusCode int           `json:"status_code"`
	Data       []*bills.Bill `json:"data"`
}

func NewGetTaxesForm() *GetTaxesForm {
	return &GetTaxesForm{}
}

func (form *GetTaxesForm) ValidateRequest(ctx context.Context, r *http.Request) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	if r.Method != http.MethodGet {
		form.Response.StatusCode = http.StatusMethodNotAllowed
		return fmt.Errorf("Invalid method")
	}

	return nil
}

func (form *GetTaxesForm) GetTaxesService(ctx context.Context) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	var err error

	err = form.GetBill(ctx)
	if err != nil {
		return err
	}

	err = form.GetParameter(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (form *GetTaxesForm) GetBill(ctx context.Context) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	list, err := bills.GetAllBill(ctx, database.PostgresDB.DBConnection)
	if err != nil {
		return err
	}

	if len(list) <= 0 {
		return fmt.Errorf("List is empty")
	}

	form.BillList = list

	return nil
}

func (form *GetTaxesForm) GetParameter(ctx context.Context) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	for i := range form.BillList {
		err := getTaxParams(form.BillList[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func getTaxParams(bill *bills.Bill) error {
	var tax float64
	var isRefundable bool
	switch bill.Code {
	case state.TaxCodeFood:
		tax = bill.Price * 0.1
		isRefundable = true
		break
	case state.TaxCodeTobacco:
		tax = (bill.Price * 0.02) + 10
		isRefundable = false
		break
	case state.TaxCodeEntertainment:
		tax = 0
		if bill.Price >= 100 {
			tax = (bill.Price - 100) * 0.01
		}
		isRefundable = false
		break
	default:
		return fmt.Errorf("Tax code '%d' is not found", bill.Code)
	}

	bill.IsRefundable = isRefundable
	bill.Tax = tax
	bill.Amount = bill.Price + tax

	return nil
}
