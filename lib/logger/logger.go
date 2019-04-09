package logger

import (
	"log"
)

const (
	// LevelDebug :
	LevelDebug = "DEBUG"
	// LevelInfo  :
	LevelInfo = "INFO"
	// LevelError :
	LevelError = "ERROR"
)

var logger Logger

// Debug ...
func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

// Info ...
func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

// Error ...
func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

// Logger ...
type Logger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// NewStandardLogger ...
func NewStandardLogger() *StandardLogger {
	return new(StandardLogger)
}

// StandardLogger ...
type StandardLogger struct {
}

// Info ...
func (StandardLogger) Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

// Debug ...
func (StandardLogger) Debug(msg string, args ...interface{}) {
	log.Printf("[DEBUG] "+msg, args...)
}

// Error ...
func (StandardLogger) Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}
