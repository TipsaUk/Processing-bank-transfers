package main

import (
	"log"
	"processing-bank-transfers/internal/config"
	"processing-bank-transfers/internal/db"
	"processing-bank-transfers/internal/migration"
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
