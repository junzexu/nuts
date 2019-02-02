package logging

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

// %(name)s	Name of the logger (logging channel)
// %(levelno)s	Numeric logging level for the message (DEBUG, INFO, WARN, ERROR, CRITICAL)
// %(levelname)s	Text logging level for the message ("DEBUG", "INFO", "WARN", "ERROR", "CRITICAL")
// %(pathname)s	Full pathname of the source file where the logging call was issued (if available)
// %(filename)s	Filename portion of pathname
// %(module)s	Module from which logging call was made
// %(lineno)d	Source line number where the logging call was issued (if available)
// %(created)f	Time when the LogRecord was created (time.time() return value)
// %(asctime)s	Textual time when the LogRecord was created
// %(msecs)d	Millisecond portion of the creation time
// %(relativeCreated)d	Time in milliseconds when the LogRecord was created, relative to the time the logging module was loaded (typically at application startup time)
// %(thread)d	Thread ID (if available)
// %(message)s	The result of record.getMessage(), computed just as the record is emitted

var timeStart = time.Now()
var pt = regexp.MustCompile("(\\([a-zA-Z]+\\))")

type fieldExtractor func(*LogRecord) interface{}

func nameExtractor(r *LogRecord) interface{} {
	return r.Name
}

func levelnoExtractor(r *LogRecord) interface{} {
	return r.Level
}

func levelnameExtractor(r *LogRecord) interface{} {
	return GetLevelName(r.Level)
}

func pathnameExtractor(r *LogRecord) interface{} {
	f, _ := runtime.FuncForPC(r.ProgramCouter).FileLine(r.ProgramCouter)
	return f
}

func filenameExtractor(r *LogRecord) interface{} {
	f, _ := runtime.FuncForPC(r.ProgramCouter).FileLine(r.ProgramCouter)
	return filepath.Base(f)
}

func linenoExtractor(r *LogRecord) interface{} {
	_, line := runtime.FuncForPC(r.ProgramCouter).FileLine(r.ProgramCouter)
	return line
}

func createdExtractor(r *LogRecord) interface{} {
	return float64(r.Timestamp.UnixNano()) / 1000000000.0
}

func asctimeExtractor(r *LogRecord) interface{} {
	return r.Timestamp.String()
}

func msecsExtractor(r *LogRecord) interface{} {
	return r.Timestamp.Nanosecond() / 1000000
}

func relativeCreatedExtractor(r *LogRecord) interface{} {
	return r.Timestamp.Sub(timeStart).Nanoseconds() / 1000000
}

func threadExtractor(r *LogRecord) interface{} {
	return 0
}

func messageExtractor(r *LogRecord) interface{} {
	return fmt.Sprintf(r.MessageFormat, r.Args...)
}

// DefaultFormatter .
type DefaultFormatter struct {
	OriFormatStr string
	FormatStr    string
	extractors   []fieldExtractor
}

// NewFormatter .
func NewFormatter(f string) Formatter {
	fmter := &DefaultFormatter{
		OriFormatStr: f,
		extractors:   []fieldExtractor{},
		FormatStr:    "",
	}

	parts := pt.FindAllStringSubmatch(f, -1)
	for _, p := range parts {
		switch p[1] {
		case "(name)":
			fmter.extractors = append(fmter.extractors, nameExtractor)
		case "(levelno)":
			fmter.extractors = append(fmter.extractors, levelnameExtractor)
		case "(levelname)":
			fmter.extractors = append(fmter.extractors, levelnameExtractor)
		case "(pathname)":
			fmter.extractors = append(fmter.extractors, pathnameExtractor)
		case "(filename)":
			fmter.extractors = append(fmter.extractors, filenameExtractor)
		case "(lineno)":
			fmter.extractors = append(fmter.extractors, linenoExtractor)
		case "(created)":
			fmter.extractors = append(fmter.extractors, createdExtractor)
		case "(asctime)":
			fmter.extractors = append(fmter.extractors, asctimeExtractor)
		case "(msecs)":
			fmter.extractors = append(fmter.extractors, msecsExtractor)
		case "(relativeCreated)":
			fmter.extractors = append(fmter.extractors, relativeCreatedExtractor)
		case "(thread)":
			fmter.extractors = append(fmter.extractors, threadExtractor)
		case "(message)":
			fmter.extractors = append(fmter.extractors, messageExtractor)
		default:
			continue
		}
	}

	fmter.FormatStr = string(pt.ReplaceAll([]byte(f), []byte("")))
	return fmter
}

// Format .
func (f *DefaultFormatter) Format(r *LogRecord) string {
	pts := make([]interface{}, 0, len(f.extractors))
	for _, p := range f.extractors {
		pts = append(pts, p(r))
	}

	return fmt.Sprintf(f.FormatStr, pts...)
}
