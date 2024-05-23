package webhook_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cfindlayisme/wmb/database"
	"github.com/cfindlayisme/wmb/webhook"
	"github.com/stretchr/testify/require"
)

func TestSubscribePrivmsg(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "An error was not expected when opening a stub database connection")
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS PrivmsgSubscriptions \\(Target TEXT, URL TEXT, FailureCount INTEGER DEFAULT 0, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY \\(Target, URL\\)\\)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()

	// Expect the SELECT query
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectPrepare("INSERT OR IGNORE INTO PrivmsgSubscriptions")
	mock.ExpectExec("INSERT OR IGNORE INTO PrivmsgSubscriptions").WithArgs("target", "url").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function
	webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
}

func TestSubscribePrivmsgExistsAndDoesnt(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "An error was not expected when opening a stub database connection")
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations for the first call
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS PrivmsgSubscriptions \\(Target TEXT, URL TEXT, FailureCount INTEGER DEFAULT 0, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY \\(Target, URL\\)\\)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()

	// Expect the SELECT query
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Call the function
	result := webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.False(t, result, "Expected false")

	// Expectations for the second call
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS PrivmsgSubscriptions \\(Target TEXT, URL TEXT, FailureCount INTEGER DEFAULT 0, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY \\(Target, URL\\)\\)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()

	// Expect the SELECT query
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectPrepare("INSERT OR IGNORE INTO PrivmsgSubscriptions")
	mock.ExpectExec("INSERT OR IGNORE INTO PrivmsgSubscriptions").WithArgs("target", "url").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function
	result = webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.True(t, result, "Expected true")
}
