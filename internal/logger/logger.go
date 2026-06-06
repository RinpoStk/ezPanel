package logger

import (
	"fmt"
	"log"
	"os"
)

var std = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

// Info logs an informational message.
func Info(format string, args ...any) {
	std.Output(2, fmt.Sprintf("[INFO] "+format, args...))
}

// Warn logs a warning message.
func Warn(format string, args ...any) {
	std.Output(2, fmt.Sprintf("[WARN] "+format, args...))
}

// Error logs an error message.
func Error(format string, args ...any) {
	std.Output(2, fmt.Sprintf("[ERROR] "+format, args...))
}

// Fatal logs an error message and exits the process.
func Fatal(format string, args ...any) {
	std.Output(2, fmt.Sprintf("[FATAL] "+format, args...))
	os.Exit(1)
}
