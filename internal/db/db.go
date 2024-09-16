package db

import (
    "database/sql"
	// "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Initialize the database connection once
func InitDB(dsn string) error {
    var err error
    DB, err = sql.Open("mysql", dsn)
	return err
    // if err != nil {
    //     log.Fatalf("Error opening database: %v", err)
    // }

    // // Test the connection
    // if err = db.Ping(); err != nil {
    //     log.Fatalf("Error connecting to the database: %v", err)
    // }
}