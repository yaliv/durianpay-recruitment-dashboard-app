package start

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func seedUsers(db *sqlx.DB) error {
	var cnt int
	row := db.QueryRow("SELECT COUNT(1) FROM users")
	if err := row.Scan(&cnt); err != nil {
		return err
	}
	if cnt == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?);", "cs@test.com", string(hash), "cs"); err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?);", "operation@test.com", string(hash), "operation"); err != nil {
			return err
		}
	}

	return nil
}

func seedPayments(db *sqlx.DB) error {
	var cnt int
	row := db.QueryRow("SELECT COUNT(1) FROM payments")
	if err := row.Scan(&cnt); err != nil {
		return err
	}
	if cnt == 0 {
		if _, err := db.Exec(`INSERT INTO payments (merchant_name, date, amount, status) VALUES
			('Amazon', '2026-01-05 10:15:00', 120000, 'completed'),
			('Netflix', '2026-01-07 20:30:00', 54000, 'completed'),
			('Spotify', '2026-01-08 09:00:00', 49999, 'completed'),
			('Steam', '2026-01-10 18:22:00', 150000, 'processing'),
			('Google Play', '2026-01-11 12:45:00', 75000, 'completed'),
			('Apple App Store', '2026-01-12 14:10:00', 99000, 'failed'),
			('Tokopedia', '2026-01-13 16:05:00', 210000, 'completed'),
			('Shopee', '2026-01-14 11:20:00', 89000, 'completed'),
			('Gojek', '2026-01-15 08:55:00', 32000, 'completed'),
			('Grab', '2026-01-16 19:40:00', 45000, 'processing'),
			('PLN', '2026-01-17 07:30:00', 350000, 'completed'),
			('Telkomsel', '2026-01-18 13:25:00', 100000, 'completed'),
			('Indihome', '2026-01-19 21:10:00', 375000, 'failed'),
			('Alfamart', '2026-01-20 17:50:00', 67000, 'completed'),
			('Indomaret', '2026-01-21 15:05:00', 82000, 'completed')
		;`); err != nil {
			return err
		}
	}

	return nil
}
