package logger

import (
	"errors"
	"log/slog"
	"os"
)

func checkLogger() {
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
	}
}

func initLogLevel(level LogLevel) error {
	switch level {
	case LogNone, LogDebug, LogInfo, LogWarn, LogError:
		logLevel = level
		checkLogger()
		return nil
	default:
		return errors.New("invalid log level")
	}
}
