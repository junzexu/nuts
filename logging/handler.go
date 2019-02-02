package logging

import (
	"io"
	"os"
	"strings"
)

// DefaultConsoleHandler .
var DefaultConsoleHandler = &ConsoleHandler{
	Formatter: NewFormatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s"),
	Stream:    os.Stdout,
	Level:     LevelDebug,
	Filters:   nil,
}

// ConsoleHandler .
type ConsoleHandler struct {
	Formatter Formatter
	Stream    io.WriteCloser
	Level     LogLevel
	Filters   []Filter
}

// SetLevel .
func (c *ConsoleHandler) SetLevel(lvl LogLevel) {
	c.Level = lvl
}

// SetFormatter .
func (c *ConsoleHandler) SetFormatter(ft Formatter) {
	c.Formatter = ft
}

// AddFilter .
func (c *ConsoleHandler) AddFilter(fl Filter) {
	c.Filters = append(c.Filters, fl)
}

// WriteRecord .
func (c *ConsoleHandler) WriteRecord(r *LogRecord) {
	str := c.Formatter.Format(r)
	str = strings.TrimRight(str, "\r\n")
	str = fontColored(r.Level, str)
	str = str + "\r\n"
	c.Stream.Write([]byte(str))
}

func fontColored(level LogLevel, str string) string {
	var color string
	switch level {
	case LevelDebug:
		color = "0;34"
	case LevelInfo:
		color = "0;32"
	case LevelWarn:
		color = "0;36"
	case LevelError:
		color = "0;31"
	case LevelCritical:
		color = "0;35"
	default:
		return str
	}

	return "\x1b[" + color + "m" + str + "\x1b[m"
}
