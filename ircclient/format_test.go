package ircclient_test

import (
	"fmt"
	"testing"

	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/cfindlayisme/wmb/model"
	"github.com/stretchr/testify/require"
)

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

func TestFormatMessage(t *testing.T) {
	// Test case for each colour code
	for i := 0; i <= 15; i++ {
		iPointer := int8(i)
		msg := model.IncomingMessage{
			Message:    "Hello World",
			ColourCode: &iPointer,
		}
		expected := fmt.Sprintf("\x03%02dHello World\x03", i)
		result := ircclient.FormatMessage(msg)
		require.Equal(t, expected, result, fmt.Sprintf("The message should be formatted with colour code %d", i))
	}

	// Test case with no colour code
	msg := model.IncomingMessage{
		Message: "Hello World",
	}
	expected := "Hello World"
	result := ircclient.FormatMessage(msg)
	require.Equal(t, expected, result, "The message should be unchanged")
}
