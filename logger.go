package log

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Level type
type Level uint32

const (
	// ErrorLevel level. Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel Level = iota
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the application.
	InfoLevel
	// VerboseLevel level. Usually only enabled when debugging. Very verbose logging.
	VerboseLevel
	// DebugLevel level. Usually only enabled when debugging. Even more verbose than verbose logging.
	DebugLevel
)

var (
	defaultLevel Level = InfoLevel
)

func levelToString(lvl Level) (string, error) {
	switch lvl {
	case DebugLevel:
		return "debug", nil
	case VerboseLevel:
		return "verbose", nil
	case InfoLevel:
		return "info", nil
	case WarnLevel:
		return "warn", nil
	case ErrorLevel:
		return "error", nil
	}

	return "", fmt.Errorf("not a valid log Level: %q", lvl)
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "verbose":
		return VerboseLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid log Level: %q", lvl)
}

// Logger represents the logger type.
type Logger struct {
	// List of transports.
	transports []Transport

	// The logging level the logger should log at.
	Level Level

	// Used for syncing transport writing.
	mu sync.Mutex

	// Reusable empty entry for managing multiple concurrent items.
	entryPool sync.Pool

	// Default log fields.
	Defaults []*FieldPair
}

// NewLogger returns a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Level:    defaultLevel,
		Defaults: make([]*FieldPair, 0),
	}
}

// WithLevel Sets the log level for logger
func (l *Logger) WithLevel(lvl Level) *Logger {
	l.Level = lvl
	return l
}

// WithTransports takes a variadic number of transports and adds them to the logger tranpsorts
func (l *Logger) WithTransports(args ...Transport) *Logger {
	l.transports = append(l.transports, args...)
	return l
}

func isZeroString(s string) bool {
	return s == ""
}

func isZeroInt(i int) bool {
	return i == 0
}

func isZeroFloat(i float64) bool {
	return i == 0
}

// Info logs message at info level.
func (l *Logger) Info(msg string) {
	l.Log(InfoLevel, msg)
}

// Error logs message at error level.
func (l *Logger) Error(msg string) {
	l.Log(ErrorLevel, msg)
}

// Warn logs message at warn level.
func (l *Logger) Warn(msg string) {
	l.Log(WarnLevel, msg)
}

// Debug logs message at debug level.
func (l *Logger) Debug(msg string) {
	l.Log(DebugLevel, msg)
}

// Verbose logs message at verbose level.
func (l *Logger) Verbose(msg string) {
	l.Log(VerboseLevel, msg)
}

// newEntry attempts to get an entry from the pool or creates a new one
func (l *Logger) newEntry() *Entry {
	entry, ok := l.entryPool.Get().(*Entry)

	if ok {
		return entry
	}

	return NewEntry(l)
}

// Log creates a new entry and calls log with level and fields
func (l *Logger) Log(level Level, msg string) {
	entry := l.newEntry()
	entry.Log(level, msg)
}

// releaseEntry adds entry back in
func (l *Logger) releaseEntry(entry *Entry) {
	// Reset the data fields
	entry.Data = &Fields{}
	entry.Time = time.Time{} // Invoking an empty time.Time struct literal will return Go's zero date.
	l.entryPool.Put(entry)
}

// WithDefaults set the default fields for a log entry
func (l *Logger) WithDefaults(args ...*FieldPair) *Logger {
	for _, arg := range args {
		// Check if default exists
		exists := false
		for _, f := range l.Defaults {
			if arg.Name == f.Name {
				exists = true
				// Override the existing default value
				f.Value = arg.Value
				break
			}
		}

		if !exists {
			l.Defaults = append(l.Defaults, arg)
		}
	}

	return l
}

// Format transforms entry data into a byte array. Can be
// extended later to include different formatter types
func (l *Logger) Format(entry *Entry) ([]byte, error) {
	entry.Data.Timestamp = entry.Time.Format(time.RFC3339)
	sev, err := levelToString(entry.Level)
	if err != nil {
		return nil, err
	}
	entry.Data.Severity = sev
	return json.Marshal(entry.Data)
}

// WithFields takes field pairs and reutrns an entry
func (l *Logger) WithFields(args ...*FieldPair) *Entry {
	entry := l.newEntry()
	return entry.WithFields(args...)
}
