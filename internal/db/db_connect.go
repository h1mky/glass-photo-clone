package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var DB *sqlx.DB

func dbConnect() {
	connect := os.Getenv("DB_CONNECT")
	dsn := fmt.Sprintf("%s", connect)

	var err error

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

}
