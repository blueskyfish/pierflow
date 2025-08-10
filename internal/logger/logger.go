package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var logLevel = LogInfo
var logger *slog.Logger

func InitLogLevel(level LogLevel) error {
	return initLogLevel(level)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Info(msg string) {
	logger.Info(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Fatal(msg string) {
	logger.Error(msg)
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func DebugWithContext(ctx context.Context, msg string) {
	logger.DebugContext(ctx, msg)
}

func InfoWithContext(ctx context.Context, msg string) {
	logger.InfoContext(ctx, msg)
}

func WarnWithContext(ctx context.Context, msg string) {
	logger.WarnContext(ctx, msg)
}

func ErrorWithContext(ctx context.Context, msg string) {
	logger.ErrorContext(ctx, msg)
}

func FatalWithContext(ctx context.Context, msg string) {
	logger.ErrorContext(ctx, msg)
	os.Exit(1)
}

func DebugfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.DebugContext(ctx, fmt.Sprintf(format, args...))
}

func InfofWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.InfoContext(ctx, fmt.Sprintf(format, args...))
}

func WarnfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.WarnContext(ctx, fmt.Sprintf(format, args...))
}

func ErrorfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.ErrorContext(ctx, fmt.Sprintf(format, args...))
}

func FatalfWithContext(ctx context.Context, format string, args ...interface{}) {
	logger.ErrorContext(ctx, fmt.Sprintf(format, args...))
	os.Exit(1)
}
