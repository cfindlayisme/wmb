package database

import (
	"database/sql"
	"os"
	"testing"
)

func TestOpenAndClose(t *testing.T) {
	// Setup
	testDBFile := "./test.db"
	defer os.Remove(testDBFile) // clean up file after test

	// Test Open
	err := DB.Open(testDBFile)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	// Test GetDB
	db := DB.GetDB()
	if db == nil {
		t.Fatalf("Failed to get database")
	}

	// Test table creation
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='PrivmsgSubscriptions'")
	var name string
	err = row.Scan(&name)
	if err != nil {
		t.Fatalf("Failed to scan row: %v", err)
	}
	if name != "PrivmsgSubscriptions" {
		t.Fatalf("Failed to create table: expected 'PrivmsgSubscriptions', got '%s'", name)
	}
}

func TestSetDB(t *testing.T) {
	// Setup
	testDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Test SetDB
	DB.SetDB(testDB)
	db := DB.GetDB()
	if db != testDB {
		t.Fatalf("Failed to set database")
	}
}
