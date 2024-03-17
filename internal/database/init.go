package database

import (
	"database/sql"
	"fmt"
	"log"
	"server/internal/config"
	"server/internal/services"
)

func InitDb(Db *sql.DB) {
	// Create products table if it doesn't exist.
	_, err := Db.Exec(`CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        created_at DATETIME NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create user table if it doesn't exist.
	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        token TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Checking of existing "admin"
	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", config.GetEnv(config.SuperAdminUsername)).Scan(&count)
	if err != nil {
		panic(err)
	}

	// Add admin in DB
	if count == 0 {
		hashPassword := services.GetMD5Hash(config.GetEnv(config.SuperAdminPassword))
		username := config.GetEnv(config.SuperAdminUsername)
		tokenStr, err := services.GenerateToken(username, hashPassword)
		if err != nil {
			panic(err)
		}

		_, err = Db.Exec(
			"INSERT INTO users (username, password, token) VALUES (?, ?, ?)",
			username,
			hashPassword,
			tokenStr,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("Admin was added")
	} else {
		fmt.Println("Admin exist")
	}

	//
}
