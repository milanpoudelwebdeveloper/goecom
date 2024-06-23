package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/milanpoudelwebdeveloper/goecom/cmd/api"
	"github.com/milanpoudelwebdeveloper/goecom/config"
	"github.com/milanpoudelwebdeveloper/goecom/db"
)

func main() {
	cfg := config.Envs

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBHost,
	)
	db, err := db.NewMyPSQLStorage(connStr)
	initStorage(db)
	if err != nil {
		log.Fatal("Sonething went wrong while connecting to database", err)

	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Something went wrong while running server", err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Something wnt wrong while connecting and pinging to database", err)
	}
	log.Println("Database is successfully connected")
}
