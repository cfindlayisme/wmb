package webhook

import (
	"log"

	"github.com/cfindlayisme/wmb/database"
	_ "github.com/mattn/go-sqlite3"
)

func SendPrivmsgWebhook(target string, message string) {
	db := database.DB.GetDB()

	// Prepare the query
	stmt, err := db.Prepare("SELECT URL FROM PrivmsgSubscriptions WHERE Target = ?")
	if err != nil {
		log.Fatalf("Error preparing the query: %v", err)
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(target)
	if err != nil {
		log.Fatalf("Error executing the query: %v", err)
	}
	defer rows.Close()

	// Loop through the rows
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			log.Fatalf("Error scanning the row: %v", err)
		}

		log.Printf("Sending webhook about message in %s to %s\n", target, url)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}
}
