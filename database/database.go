package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

var DB = &Database{}

func (d *Database) Open(databaseFile string) {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatalf("Error opening SQLite database: %v", err)
	}
	d.db = db
}

func (d *Database) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) SetDB(db *sql.DB) {
	d.db = db
}
