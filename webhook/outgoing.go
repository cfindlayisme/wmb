package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cfindlayisme/wmb/database"
	"github.com/cfindlayisme/wmb/model"
	_ "github.com/mattn/go-sqlite3"
)

func SendPrivmsgWebhook(target string, message string) {
	db := database.DB.GetDB()

	// Prepare the query
	stmt, err := db.Prepare("SELECT URL, FailureCount FROM PrivmsgSubscriptions WHERE Target = ?")
	if err != nil {
		log.Fatalf("Error preparing the query: %v", err)
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(target)
	if err != nil {
		log.Fatalf("Error executing the query: %v", err)
	}

	// Create a slice to hold the URLs and failure counts
	var urls []string
	var failureCounts []int

	// Loop through the rows
	for rows.Next() {
		var url string
		var failureCount int
		if err := rows.Scan(&url, &failureCount); err != nil {
			log.Fatalf("Error scanning the row: %v", err)
		}

		urls = append(urls, url)
		failureCounts = append(failureCounts, failureCount)
	}

	// Print the number of rows returned by the query
	log.Println("Number of rows returned by the query:", len(urls))

	// Close the rows before sending the webhooks
	rows.Close()

	// Send the webhooks
	for i, url := range urls {
		msg := model.DirectedOutgoingMessage{
			Target:  target,
			Message: message,
		}

		err := sendPrivmsgWebhookToUrl(url, msg)
		if err != nil {
			log.Println("Failed to send to the target URL:", url)
			updateFailureCount(target, url)
			failureCounts[i]++
			if failureCounts[i] >= 3 {
				log.Println("Failed to send to the target URL:", url)
			}
		}
	}

	// Print the number of webhooks sent
	log.Println("Number of webhooks sent:", len(urls))

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v", err)
	}
}

func updateFailureCount(target string, url string) {
	db := database.DB.GetDB()

	// Prepare the update statement for failure
	updateFailureStmt, err := db.Prepare("UPDATE PrivmsgSubscriptions SET FailureCount = FailureCount + 1 WHERE Target = ? AND URL = ?")
	if err != nil {
		log.Fatalf("Error preparing the update statement: %v", err)
	}
	defer updateFailureStmt.Close()

	_, err = updateFailureStmt.Exec(target, url)
	if err != nil {
		log.Fatalf("Error updating the failure count: %v", err)
	}
}

func sendPrivmsgWebhookToUrl(url string, msg model.DirectedOutgoingMessage) error {
	// Convert the message to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	log.Println("Sending message to", url)
	log.Println(msg)

	// Send the request
	client := &http.Client{
		Timeout: time.Second * 1, // Set the timeout to 1 second
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
