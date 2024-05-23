package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

var DB = &Database{}

func (d *Database) Open(databaseFile string) error {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) Close() error {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			return err
		}
		d.db = nil
	}
	return nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) SetDB(db *sql.DB) {
	d.db = db
}
