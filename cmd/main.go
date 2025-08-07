package main

import (
	"database/sql"
	"log"

	"github.com/diegobbrito/ecom/cmd/api"
	"github.com/diegobbrito/ecom/config"
	"github.com/diegobbrito/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Database connection established successfully")
}
