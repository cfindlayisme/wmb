package webhook

import (
	"database/sql"
	"log"

	"github.com/cfindlayisme/wmb/database"
	_ "github.com/mattn/go-sqlite3"
)

func SubscribePrivmsg(target string, url string) bool {
	db := database.DB.GetDB()

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)
		return false
	}
	defer tx.Rollback() // Rollback the transaction if it's not committed

	// Check if the data already exists
	row := tx.QueryRow("SELECT 1 FROM PrivmsgSubscriptions WHERE Target = ? AND URL = ?", target, url)
	var exists bool
	err = row.Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking existence: %v", err)
		return false
	}
	if exists {
		return false
	}

	// Prepare the statement
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO PrivmsgSubscriptions (Target, URL, FailureCount) VALUES (?, ?, 0)")
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return false
	}
	defer stmt.Close() // Close the statement when it's no longer needed

	// Execute the statement
	_, err = stmt.Exec(target, url)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return false
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return false
	}

	return true
}

func UnsubscribePrivmsg(target string, url string) bool {
	db := database.DB.GetDB()

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)
		return false
	}
	defer tx.Rollback() // Rollback the transaction if it's not committed

	// Check if the data already exists
	row := tx.QueryRow("SELECT 1 FROM PrivmsgSubscriptions WHERE Target = ? AND URL = ?", target, url)
	var exists bool
	err = row.Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking existence: %v", err)
		return false
	}
	if !exists {
		return false
	}

	// Prepare the statement
	stmt, err := tx.Prepare("DELETE FROM PrivmsgSubscriptions WHERE Target = ? AND URL = ?")
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return false
	}
	defer stmt.Close() // Close the statement when it's no longer needed

	// Execute the statement
	_, err = stmt.Exec(target, url)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return false
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return false
	}

	return true
}
