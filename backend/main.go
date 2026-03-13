package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/durianpay/fullstack-boilerplate/internal/config"
	srv "github.com/durianpay/fullstack-boilerplate/internal/service/http"
	"github.com/durianpay/fullstack-boilerplate/start"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "dashboard.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := start.InitDB(db); err != nil {
		log.Fatal(err)
	}

	apiHandler := start.SetupAPIHandler(db)

	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation)

	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}
