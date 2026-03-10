package entity

import (
	"time"
)

type Payment struct {
	Amount    string    `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	Id        string    `db:"id"`
	Merchant  string    `db:"merchant"`
	Status    string    `db:"status"`
}
