package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection() (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v %v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PARAMS"))

	log.Println("try connecting database : ", dsn)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return
	}

	log.Println("success connecting database")
	return
}
