package bills

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	utilctx "github.com/naufalridho/tax-calculator/common/context"
)

type Bill struct {
	ID           int64   `db:"id" json:"id"`
	Name         string  `db:"name" json:"name"`
	Code         int     `db:"tax_code" json:"code"`
	Price        float64 `db:"price" json:"price"`
	IsRefundable bool    `db:"-" json:"is_refundable"`
	Tax          float64 `db:"-" json:"tax"`
	Amount       float64 `db:"-" json:"amount"`
}

const getAllBill = `
		SELECT
			id, name, tax_code, price
		FROM bill
		ORDER BY id
	`
const getBillbyID = `
		SELECT
			?
		FROM bill
		ORDER BY id desc
	`
const insertBill = `
		INSERT INTO bill
			(name, tax_code, price) 
		VALUES
			(?, ?, ?)
	`

func GetAllBill(ctx context.Context, db *sqlx.DB) ([]*Bill, error) {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	bill := []*Bill{}

	query, args, err := sqlx.In(getAllBill)
	if err != nil {
		return bill, fmt.Errorf("Failed to bind query. Err:, %s", err.Error())
	}

	query = db.Rebind(query)
	err = db.Select(&bill, query, args...)
	if err != nil {
		return bill, fmt.Errorf("Failed to get bills. Err:, %s", err.Error())
	}

	return bill, nil
}

func GetBillbyID(ctx context.Context, db *sqlx.DB, id int64) (*Bill, error) {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	bill := &Bill{}

	query, args, err := sqlx.In(getBillbyID, id)
	if err != nil {
		return bill, fmt.Errorf("Failed to bind query. Err:, %s", err.Error())
	}

	query = db.Rebind(query)
	err = db.Select(bill, query, args...)
	if err != nil {
		return bill, fmt.Errorf("Failed to get bill. Err:, %s", err.Error())
	}

	return bill, nil
}

func InsertBill(ctx context.Context, db *sqlx.DB, name string, code int, price float64) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("Failed to start database transaction. Err: %s", err.Error())
	}

	query := tx.Rebind(insertBill)
	_, err = tx.Exec(query, name, code, price)
	if err != nil {
		return fmt.Errorf("Failed to insert bills. Err: %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Failed to commit database transaction. Err: %s", err.Error())
	}

	return nil
}
