package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func QueryBool(ctx echo.Context, key string, defaultValue bool) bool {
	value := ctx.QueryParam(key)
	if value == "" {
		return defaultValue
	}
	return strings.ToLower(value) == "true"
}

func HeaderUser(ctx echo.Context) string {
	return ctx.Request().Header.Get("x-pierflow-user")
}
