package env_test

import (
	"os"
	"testing"

	"github.com/cfindlayisme/wmb/env"
	"github.com/go-playground/assert/v2"
)

func TestGetServer(t *testing.T) {
	expected := "localhost"
	os.Setenv("IRC_SERVER", expected)

	result := env.GetServer()

	assert.Equal(t, result, expected)
}

func TestGetChannel(t *testing.T) {
	expected := "#wmb"
	defaultResult := "#wmb"

	os.Setenv("IRC_CHANNEL", expected)

	result := env.GetChannel()

	assert.Equal(t, result, expected)

	os.Unsetenv("IRC_CHANNEL")
	result = env.GetChannel()

	assert.Equal(t, result, defaultResult)
}

func TestGetNick(t *testing.T) {
	expected := "blah"
	defaultResult := "wmb"

	os.Setenv("IRC_NICK", expected)

	result := env.GetNick()

	assert.Equal(t, result, expected)

	os.Unsetenv("IRC_NICK")
	result = env.GetNick()

	assert.Equal(t, result, defaultResult)
}

func TestGetDatabaseFile(t *testing.T) {
	// Test default value
	os.Unsetenv("DBFILE")
	dbfile := env.GetDatabaseFile()
	if dbfile != "wmb.db" {
		t.Fatalf("Expected default dbfile to be 'wmb.db', got '%s'", dbfile)
	}

	// Test environment variable
	os.Setenv("DBFILE", "test.db")
	dbfile = env.GetDatabaseFile()
	if dbfile != "test.db" {
		t.Fatalf("Expected dbfile to be 'test.db', got '%s'", dbfile)
	}
}
