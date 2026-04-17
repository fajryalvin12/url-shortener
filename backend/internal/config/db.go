package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	dsn := os.Getenv("DB_URL")
	fmt.Println("DSN:", dsn)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic("DB Connection Error:" + err.Error())
	}

	return db
}