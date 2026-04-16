package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	dsn := "postgres://postgres:1@localhost:5432/url_shortener?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("DB Connection Error!", err)
	}

	return db
}