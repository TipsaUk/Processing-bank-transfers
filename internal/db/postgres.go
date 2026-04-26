package db

import (
	"database/sql"
	"log"
	"time"
)

func InitDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	// пул соединений
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	log.Println("DB connected")

	return db
}
