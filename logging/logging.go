package logging

import (
	"fmt"
	"time"
)

// [pep-0282](https://www.python.org/dev/peps/pep-0282/) implementation
// > 1. These Logger objects create LogRecord objects which are passed to Handler objects for output.
// > 2. Both Loggers and Handlers may use logging levels and (optionally) Filters to decide if they are interested in a particular LogRecord.
// > 3. When it is necessary to output a LogRecord externally, a Handler can (optionally) use a Formatter to localize and format the message before sending it to an I/O stream.

// LogLevel .
type LogLevel uint

// Logger .
// Each Logger keeps track of a set of output Handlers.
// By default all Loggers also send their output to all Handlers of their ancestor Loggers.
// Loggers may, however, also be configured to ignore Handlers higher up the tree.
type Logger interface {
	// Log .
	AddHandler(hl Handler)

	SetLevel(lvl LogLevel)
	// Return true if requests at level 'lvl' will NOT be discarded.
	IsEnabledFor(lvl LogLevel) bool
	// If a logger's level is not set, the system consults all its ancestors, walking up the hierarchy until an explicitly set level is found.
	// That is regarded as the "effective level" of the logger.
	GetEffectiveLevel() LogLevel
	SetRollover(maxBytes uint, backupCount uint)
	Close()

	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Critical(msg string, args ...interface{})
	Exception(msg string, args ...interface{})
}

// LoggerFactory .
type LoggerFactory interface {
	GetLogger(name string) Logger
	GetParent(name string) Logger
}

// Handler .
type Handler interface {
	SetFormatter(ft Formatter)
	AddFilter(fl Filter)
	SetLevel(lvl LogLevel)

	WriteRecord(r *LogRecord)
}

// Filter .
type Filter interface {
	// Return a value indicating true if the record is to be
	// processed.  Possibly modify the record, if deemed
	// appropriate by the filter.
	Filte(r *LogRecord) bool
}

// Formatter .
type Formatter interface {
	Format(r *LogRecord) string
}

// LogRecord .
type LogRecord struct {
	Name          string
	ProgramCouter uintptr
	Level         LogLevel
	Timestamp     time.Time
	MessageFormat string
	Args          []interface{}
}

// GetMessage .
func (l *LogRecord) GetMessage() Message {
	return nil
}

// Message .
type Message interface {
	fmt.Stringer
	MessgeID() string
}
