package logging

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func captureLog(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	f()
	return buf.String()
}

// Debug only fires when DEBUG=true

func TestDebugSilentWhenDebugOff(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Debug("should not appear") })
	assert.Equal(t, out, "")
}

func TestDebugOutputWhenDebugOn(t *testing.T) {
	os.Setenv("DEBUG", "true")
	defer os.Unsetenv("DEBUG")
	out := captureLog(func() { Debug("test message") })
	assert.Equal(t, true, len(out) > 0)
	assert.Equal(t, true, containsSubstr(out, "[DEBUG]"))
	assert.Equal(t, true, containsSubstr(out, "test message"))
}

func TestDebugfSilentWhenDebugOff(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Debugf("should not appear %s", "value") })
	assert.Equal(t, out, "")
}

func TestDebugfOutputWhenDebugOn(t *testing.T) {
	os.Setenv("DEBUG", "true")
	defer os.Unsetenv("DEBUG")
	out := captureLog(func() { Debugf("formatted %s", "value") })
	assert.Equal(t, true, containsSubstr(out, "[DEBUG]"))
	assert.Equal(t, true, containsSubstr(out, "formatted value"))
}

// Info always fires

func TestInfoAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Info("info message") })
	assert.Equal(t, true, containsSubstr(out, "[INFO]"))
	assert.Equal(t, true, containsSubstr(out, "info message"))
}

func TestInfofAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Infof("info %s", "formatted") })
	assert.Equal(t, true, containsSubstr(out, "[INFO]"))
	assert.Equal(t, true, containsSubstr(out, "info formatted"))
}

// Warn always fires

func TestWarnAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Warn("warn message") })
	assert.Equal(t, true, containsSubstr(out, "[WARN]"))
	assert.Equal(t, true, containsSubstr(out, "warn message"))
}

func TestWarnfAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Warnf("warn %s", "formatted") })
	assert.Equal(t, true, containsSubstr(out, "[WARN]"))
	assert.Equal(t, true, containsSubstr(out, "warn formatted"))
}

// Error always fires

func TestErrorAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Error("error message") })
	assert.Equal(t, true, containsSubstr(out, "[ERROR]"))
	assert.Equal(t, true, containsSubstr(out, "error message"))
}

func TestErrorfAlwaysLogs(t *testing.T) {
	os.Unsetenv("DEBUG")
	out := captureLog(func() { Errorf("error %s", "formatted") })
	assert.Equal(t, true, containsSubstr(out, "[ERROR]"))
	assert.Equal(t, true, containsSubstr(out, "error formatted"))
}

// Color codes are present in output

func TestInfoContainsColorCodes(t *testing.T) {
	out := captureLog(func() { Info("colored") })
	assert.Equal(t, true, containsSubstr(out, colorGreen))
	assert.Equal(t, true, containsSubstr(out, colorReset))
}

func TestWarnContainsColorCodes(t *testing.T) {
	out := captureLog(func() { Warn("colored") })
	assert.Equal(t, true, containsSubstr(out, colorYellow))
	assert.Equal(t, true, containsSubstr(out, colorReset))
}

func TestErrorContainsColorCodes(t *testing.T) {
	out := captureLog(func() { Error("colored") })
	assert.Equal(t, true, containsSubstr(out, colorRed))
	assert.Equal(t, true, containsSubstr(out, colorReset))
}

func TestDebugContainsColorCodesWhenDebugOn(t *testing.T) {
	os.Setenv("DEBUG", "true")
	defer os.Unsetenv("DEBUG")
	out := captureLog(func() { Debug("colored") })
	assert.Equal(t, true, containsSubstr(out, colorCyan))
	assert.Equal(t, true, containsSubstr(out, colorReset))
}

func containsSubstr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		func() bool {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}())
}
