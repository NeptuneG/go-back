package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	dbConnOnce sync.Once
	dbConn     *sql.DB
)

func ConnectDatabase() *sql.DB {
	dbConnOnce.Do(func() {
		conn, err := sql.Open("postgres", dataSourceName())
		if err != nil {
			panic(err)
		}
		dbConn = conn
	})
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

func ExecTx(ctx context.Context, fn func(*sql.Tx) (interface{}, error)) (interface{}, error) {
	if dbConn == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	tx, err := dbConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result, err := fn(tx)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return nil, err
	}

	return result, tx.Commit()
}
