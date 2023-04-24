package db

import (
	"database/sql"
	"log"

	"github.com/hyeong01/thesaurus/config"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func Connect() {
	var err error
	cfg := config.GetConfig()
	connStr := "user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable"
	conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Unable to ping the database:", err)
	}

	log.Println("Connected to the database")
}
