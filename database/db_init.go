package database

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"time"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://myuser:secret@localhost:63892/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
