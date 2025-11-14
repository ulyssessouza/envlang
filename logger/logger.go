package logger

import (
	"io"
	"log"
	"os"
	"sync"
)

// Level type for log levels.
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelNames = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

func (l Level) String() string {
	if l >= DEBUG && l <= ERROR {
		return levelNames[l]
	}
	return "UNKNOWN"
}

// Logger is a custom logger with levels
type Logger struct {
	m     sync.Mutex
	level Level
	out   io.Writer
	std   *log.Logger
}

// New creates a new Logger instance.
func New(out io.Writer, level Level) *Logger {
	l := &Logger{
		level: level,
		out:   out,
		std:   log.New(out, "", log.LstdFlags),
	}
	return l
}

// DefaultLogger is the default logger instance.
var DefaultLogger = New(os.Stderr, INFO)

// SetOutput sets the output writer for the default logger
func SetOutput(out io.Writer) {
	DefaultLogger.m.Lock()
	defer DefaultLogger.m.Unlock()
	DefaultLogger.out = out
	DefaultLogger.std.SetOutput(out)
}

// SetLevel sets the log level for the default logger
func SetLevel(level Level) {
	DefaultLogger.m.Lock()
	defer DefaultLogger.m.Unlock()
	DefaultLogger.level = level
}

// Debugf logs a debug message
func Debugf(format string, v ...interface{}) {
	DefaultLogger.Debugf(format, v...)
}

// Infof logs an info message
func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

// Warnf logs a warning message
func Warnf(format string, v ...interface{}) {
	DefaultLogger.Warnf(format, v...)
}

// Errorf logs an error message
func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

// Debugf logs a debug message
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.std.Printf("[DEBUG] "+format, v...)
	}
}

// Infof logs an info message
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		l.std.Printf("[INFO] "+format, v...)
	}
}

// Warnf logs a warning message
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		l.std.Printf("[WARN] "+format, v...)
	}
}

// Errorf logs an error message
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.std.Printf("[ERROR] "+format, v...)
	}
}
