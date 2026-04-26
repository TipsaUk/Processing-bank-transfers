package main

import (
	"Processing-bank-transfers/internal/config"
	"Processing-bank-transfers/internal/db"
	"Processing-bank-transfers/internal/migration"
	"log"
)

func main() {
	cfg := config.Load()

	database := db.InitDB(cfg.DBConnString())

	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Run(database, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
