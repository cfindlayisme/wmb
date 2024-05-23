package webhook_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cfindlayisme/wmb/database"
	"github.com/cfindlayisme/wmb/webhook"
)

func TestSubscribePrivmsg(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS PrivmsgSubscriptions").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT OR IGNORE INTO PrivmsgSubscriptions")
	mock.ExpectExec("INSERT OR IGNORE INTO PrivmsgSubscriptions").WithArgs("target", "url").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function
	webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
