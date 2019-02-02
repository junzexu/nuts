package logging

import (
	"runtime"
	"time"
)

// DefaultLogger .
type DefaultLogger struct {
	Name     string // logger's module name
	handlers []Handler
	creator  LoggerFactory
	popup    bool
	level    LogLevel
}

// AddHandler .
func (l *DefaultLogger) AddHandler(hl Handler) {
	l.handlers = append(l.handlers, hl)
}

// SetLevel .
func (l *DefaultLogger) SetLevel(lvl LogLevel) {
	l.level = lvl
}

// Log .
func (l *DefaultLogger) log(lvl LogLevel, msg string, args ...interface{}) {
	if l.GetEffectiveLevel() <= lvl {
		if pc, _, _, ok := runtime.Caller(2); ok {
			record := &LogRecord{
				Name:          l.Name,
				ProgramCouter: pc,
				Level:         lvl,
				Timestamp:     time.Now(),
				MessageFormat: msg,
				Args:          args,
			}
			for _, h := range l.handlers {
				h.WriteRecord(record)
			}
		}
	}
}

// IsEnabledFor .
func (l *DefaultLogger) IsEnabledFor(lvl LogLevel) bool {
	return l.level >= lvl
}

// GetEffectiveLevel .
func (l *DefaultLogger) GetEffectiveLevel() LogLevel {
	return l.level
}

// SetRollover .
func (l *DefaultLogger) SetRollover(maxBytes uint, backupCount uint) {

}

// Debug .
func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
	l.log(LevelDebug, msg, args...)
}

// Info .
func (l *DefaultLogger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

// Warn .
func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
	l.log(LevelWarn, msg, args...)
}

// Error .
func (l *DefaultLogger) Error(msg string, args ...interface{}) {
	l.log(LevelError, msg, args...)
}

// Critical .
func (l *DefaultLogger) Critical(msg string, args ...interface{}) {
	l.log(LevelCritical, msg, args...)
}

// Exception .
func (l *DefaultLogger) Exception(msg string, args ...interface{}) {
	l.log(LevelCritical, msg, args...)
}

// Close .
func (l *DefaultLogger) Close() {

}
