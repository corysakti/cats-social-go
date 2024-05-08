package database

import (
	"database/sql"
	"github.com/corysakti/cats-social-go/helper"
	_ "github.com/lib/pq" // PostgreSQL driver
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://myuser:secret@localhost:50420/mydatabase?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
