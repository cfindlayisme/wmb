package webhook

import (
	"database/sql"
	"log"

	"github.com/cfindlayisme/wmb/env"
)

func SubscribePrivmsg(target string, url string) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", env.GetDatabaseFile())
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS PrivmsgSubscriptions (Target TEXT PRIMARY KEY, URL TEXT, FailureCount INTEGER DEFAULT 0, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error beginning transaction: %v", err)
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO PrivmsgSubscriptions (Target, URL, FailureCount) VALUES (?, ?, 0)")
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}

	// Execute the statement
	_, err = stmt.Exec(target, url)
	if err != nil {
		log.Fatalf("Error executing statement: %v", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}
}
