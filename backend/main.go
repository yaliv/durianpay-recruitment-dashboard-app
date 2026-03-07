package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/config"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ar "github.com/durianpay/fullstack-boilerplate/internal/module/auth/repository"
	au "github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	srv "github.com/durianpay/fullstack-boilerplate/internal/service/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	_ = godotenv.Load()

	db, err := sql.Open("sqlite3", "dashboard.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initDB(db); err != nil {
		log.Fatal(err)
	}

	JwtExpiredDuration, err := time.ParseDuration(config.JwtExpired)
	if err != nil {
		panic(err)
	}

	userRepo := ar.NewUserRepo(db)

	authUC := au.NewAuthUsecase(userRepo, config.JwtSecret, JwtExpiredDuration)

	authH := ah.NewAuthHandler(authUC)

	apiHandler := &api.APIHandler{
		Auth: authH,
	}

	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation)

	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}

func initDB(db *sql.DB) error {
	// create tables if not exist
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			role TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS payments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			merchant_name TEXT NOT NULL UNIQUE,
			date DATETIME NOT NULL,
			amount INTEGER NOT NULL,
			status TEXT NOT NULL
		);`,
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

func seedUsers(db *sql.DB) error {
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
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "cs@test.com", string(hash), "cs"); err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "operation@test.com", string(hash), "operation"); err != nil {
			return err
		}
	}

	return nil
}

func seedPayments(db *sql.DB) error {
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
			('Indomaret', '2026-01-21 15:05:00', 82000, 'completed');
		`); err != nil {
			return err
		}
	}

	return nil
}
