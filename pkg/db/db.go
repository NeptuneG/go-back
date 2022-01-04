package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
)

func ConnectDatabase() *sql.DB {
	dbConn, err := sql.Open("postgres", dataSourceName())
	if err != nil {
		panic(err)
	}
	return dbConn
}

func dataSourceName() string {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		)
	} else {
		return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			url.QueryEscape(os.Getenv("DB_PASSWORD")),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		)
	}
}
