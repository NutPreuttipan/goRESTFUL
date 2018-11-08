package driver

import (
	"database/sql"
	"log"
	"os"
	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	// connStr := "user=postgres password=password@1 dbname=postgres sslmode=disable"
	pgURL, err := pq.ParseURL(os.Getenv("POSTGRESQL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}