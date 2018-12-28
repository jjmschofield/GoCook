package logger

import (
	"go.uber.org/zap"
	"os"
)

var singleton *zap.Logger

func Init(profile string) *zap.Logger {
	var logger *zap.Logger

	if profile == "dev" {
		logger, _ = zap.NewDevelopment(zap.Fields(zap.Int("pid", os.Getpid())))
	} else {
		logger, _ = zap.NewProduction(zap.Fields(zap.Int("pid", os.Getpid())))
	}

	singleton = logger

	zap.ReplaceGlobals(logger)

	return logger
}

func WithContext(fields ...zap.Field) zap.Logger {
	contextLogger := singleton

	contextLogger = contextLogger.With(fields...)

	return *contextLogger
}

// Debug logs a debug message with the given fields
func Debug(message string, fields ...zap.Field) {
	singleton.Debug(message, fields...)
}

// Info logs a debug message with the given fields
func Info(message string, fields ...zap.Field) {
	singleton.Info(message, fields...)
}

// Warn logs a debug message with the given fields
func Warn(message string, fields ...zap.Field) {
	singleton.Warn(message, fields...)
}

// Error logs a debug message with the given fields
func Error(message string, fields ...zap.Field) {
	singleton.Error(message, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(message string, fields ...zap.Field) {
	singleton.Fatal(message, fields...)
}
