package ircclient_test

import (
	"bufio"
	"net"
	"strings"
	"testing"

	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/stretchr/testify/require"
)

func TestReturnPong(t *testing.T) {
	// Create a pair of connected, in-memory network connections
	conn1, conn2 := net.Pipe()

	// Call the function in a goroutine, because it will block until conn2 is read
	go ircclient.ReturnPong(conn1, "PING :tmi.twitch.tv")

	// Read the data from conn2
	reader := bufio.NewReader(conn2)
	data, err := reader.ReadString('\n')
	require.NoError(t, err, "Error reading from connection")

	// Remove the trailing newline
	data = strings.TrimSuffix(data, "\n")

	// Check if the data is as expected
	require.Equal(t, "PONG :tmi.twitch.tv\r", data, "The data should be 'PONG :tmi.twitch.tv'")
}

func TestCleanMessage(t *testing.T) {
	// Test case with newlines and carriage returns
	input := "Hello\nWorld\r\n"
	expected := "HelloWorld"
	result := ircclient.CleanMessage(input)
	require.Equal(t, expected, result, "The message should be cleaned")

	// Test case with no newlines or carriage returns
	input = "Hello World"
	expected = "Hello World"
	result = ircclient.CleanMessage(input)
	require.Equal(t, expected, result, "The message should be unchanged")
}
