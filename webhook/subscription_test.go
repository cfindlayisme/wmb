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
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectPrepare("^INSERT OR IGNORE INTO PrivmsgSubscriptions \\(Target, URL, FailureCount\\) VALUES \\(\\?, \\?, 0\\)$")
	mock.ExpectExec("^INSERT OR IGNORE INTO PrivmsgSubscriptions \\(Target, URL, FailureCount\\) VALUES \\(\\?, \\?, 0\\)$").
		WithArgs("target", "url").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function
	result := webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.True(t, result, "Expected true")
}

func TestSubscribePrivmsgExists(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "An error was not expected when opening a stub database connection")
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	mock.ExpectRollback()

	// Call the function
	result := webhook.SubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.False(t, result, "Expected false")
}

func TestUnsubscribePrivmsg(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "An error was not expected when opening a stub database connection")
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	mock.ExpectPrepare("^DELETE FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$")
	mock.ExpectExec("^DELETE FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function
	result := webhook.UnsubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.True(t, result, "Expected true")
}

func TestUnsubscribePrivmsgNotExists(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "An error was not expected when opening a stub database connection")
	defer db.Close()

	// Set the mock database in the DB object
	database.DB.SetDB(db)

	// Expectations
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT 1 FROM PrivmsgSubscriptions WHERE Target = \\? AND URL = \\?$").
		WithArgs("target", "url").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectRollback()

	// Call the function
	result := webhook.UnsubscribePrivmsg("target", "url")

	// Make sure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")

	// Check the result
	require.False(t, result, "Expected false")
}
