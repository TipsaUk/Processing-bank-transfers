package main

import (
	"Processing-bank-transfers/internal/config"
	"Processing-bank-transfers/internal/db"
	"Processing-bank-transfers/internal/server"
)

func main() {

	cfg := config.Load()

	dataBase := db.InitDB(cfg.DBConnString())

	srv := server.New(dataBase, cfg.ServerPort)
	srv.Start()

}
