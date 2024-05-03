package main

import (
	"log"
	"os"

	"github.com/corysakti/cats-social-go/app/config/databases"
	"github.com/corysakti/cats-social-go/wire"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error reading file env : %v", err)
		return
	}

	db, err := databases.Connection()
	if err != nil {
		log.Fatalf("error connecting database : %v", err)
		return
	}

	engine := wire.SetupApp(db)
	engine.Run(os.Getenv("APP_PORT"))

}
