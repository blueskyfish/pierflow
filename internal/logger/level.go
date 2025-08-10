package logger

import "log/slog"

type LogLevel string

const (
	LogNone  LogLevel = "none"
	LogDebug LogLevel = "debug"
	LogInfo  LogLevel = "info"
	LogWarn  LogLevel = "warn"
	LogError LogLevel = "error"
)

func (l LogLevel) Level() slog.Level {
	switch l {
	case LogNone:
		return -99
	case LogDebug:
		return slog.LevelDebug
	case LogInfo:
		return slog.LevelInfo
	case LogWarn:
		return slog.LevelWarn
	case LogError:
		return slog.LevelError
	default:
		return -99
	}
}

func (l LogLevel) String() string {
	return string(l)
}
