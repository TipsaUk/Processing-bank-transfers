package main

import (
	"processing-bank-transfers/internal/config"
	"processing-bank-transfers/internal/db"
	"processing-bank-transfers/internal/server"
)

func main() {

	cfg := config.Load()

	dataBase := db.InitDB(cfg.DBConnString())

	srv := server.New(dataBase, cfg.ServerPort)
	srv.Start()

}
