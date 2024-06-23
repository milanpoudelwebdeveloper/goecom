package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewMyPSQLStorage(cfg string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal("Something went wrong while connecting to database", err)
	}
	return db, nil

}
