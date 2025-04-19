package logging

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Logger() *slog.Logger {
	return logger
}

func Debug(message string, args ...any) {
	logger.Debug(message, args...)
}

func Info(message string, args ...any) {
	logger.Info(message, args...)
}

func Warn(message string, args ...any) {
	logger.Warn(message, args...)
}

func Error(message string, args ...any) {
	logger.Error(message, args...)
}
