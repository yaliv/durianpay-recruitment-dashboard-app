package start

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func InitDB(db *sqlx.DB) error {
	// create tables if not exist
	stmts := []string{
		createTableUsers,
		createTablePayments,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}

	// seed data if not exist
	seedUsers(db)
	seedPayments(db)

	const dbLifetime = time.Minute * 5
	db.SetConnMaxLifetime(dbLifetime)
	return nil
}
