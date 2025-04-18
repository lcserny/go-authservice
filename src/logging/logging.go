package logging

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Info(message string, args ...any) {
	if len(args) > 0 {
		logger.Info(message, args...)
	} else {
		logger.Info(message)
	}
}
