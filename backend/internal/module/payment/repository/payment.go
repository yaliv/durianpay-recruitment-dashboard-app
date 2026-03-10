package pr

import (
	"database/sql"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/helper/queryhelper"
	"github.com/jmoiron/sqlx"
)

type (
	PaymentRepository interface {
		GetPaymentByID(id string) (*entity.Payment, error)
		AllPayments(status *string, sort *string) ([]*entity.Payment, error)
		CountPayments(status *string) (int, error)
	}

	Payment struct {
		db *sqlx.DB
	}
)

func NewPaymentRepo(db *sqlx.DB) *Payment {
	return &Payment{db: db}
}

func (r *Payment) GetPaymentByID(id string) (*entity.Payment, error) {
	var p entity.Payment

	q := `SELECT id, merchant_name, date, amount, status FROM payments WHERE id = ?`
	if err := r.db.Get(&p, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrorNotFound("Record not found")
		}
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return &p, nil
}

func (r *Payment) AllPayments(status *string, sort *string) ([]*entity.Payment, error) {
	var ps []*entity.Payment

	q := `SELECT id, merchant_name, date, amount, status FROM payments`
	var params []any

	if status != nil {
		q += ` WHERE status = ?`
		params = []any{*status}
	}
	if sort != nil {
		queryhelper.AppendOrderBy(&q, *sort)
	}

	if err := r.db.Get(&ps, q, params...); err != nil {
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return ps, nil
}

func (r *Payment) CountPayments(status *string) (int, error) {
	var cnt int

	q := `SELECT COUNT(1) FROM payments`
	var params []any

	if status != nil {
		q += ` WHERE status = ?`
		params = []any{*status}
	}

	row := r.db.QueryRow(q, params...)
	if err := row.Scan(&cnt); err != nil {
		return 0, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return cnt, nil
}
