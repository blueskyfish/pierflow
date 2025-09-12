package api

import (
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func loggingRequestFunc(ctx echo.Context, v middleware.RequestLoggerValues) error {
	if v.Error != nil {
		logger.ErrorfWithContext(ctx.Request().Context(), "Request: %s %s %d (Error: %v)", v.Method, v.URI, v.Status, v.Error)
		return v.Error
	}
	logger.InfofWithContext(ctx.Request().Context(), "Request: %s %s %d (Duration: %s)", v.Method, v.URI, v.Status, v.Latency.String())
	return nil
}
