package logging

// ===== part 1: log level manage =====
const (
	// LevelDebug .
	LevelDebug LogLevel = 0
	// LevelInfo .
	LevelInfo LogLevel = 1
	// LevelWarn .
	LevelWarn LogLevel = 2
	// LevelError .
	LevelError LogLevel = 3
	// LevelCritical .
	LevelCritical LogLevel = 4
)

// log level name
var levelNames = map[LogLevel]string{
	LevelDebug:    "Debug",
	LevelInfo:     "Info",
	LevelWarn:     "Warn",
	LevelError:    "Error",
	LevelCritical: "Critical",
}

// AddLevelName .
func AddLevelName(lvl LogLevel, ln string) {
	levelNames[lvl] = ln
}

// GetLevelName .
func GetLevelName(lvl LogLevel) string {
	return levelNames[lvl]
}

var globalLevel LogLevel

// Disable .
// Do not generate any LogRecords for requests with a severity less than 'lvl'.
func Disable(lvl LogLevel) {
	globalLevel = lvl
}
