package app

import (
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"server/internal/config"
	"server/internal/database"
	"server/internal/transport/rest"
)

func InitServer() {
	// Open the database connection.
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		panic(err)
	}

	db := config.GetInstance()
	Db := db.Db
	database.InitDb(Db)
	defer db.Db.Close()

	rest.CreateRouter(Db)
}
