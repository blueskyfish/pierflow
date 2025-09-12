package projects

import (
	"github.com/blueskyfish/pierflow/internal/business/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) ProjectEventConnect(ctx echo.Context) error {
	err := pm.eventServe.Listen(ctx)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}

func (pm *ProjectManager) ProjectEventPing(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("user is required"))
	}
	err := pm.eventServe.Send(userId, "", "ping")
	if err != nil {
		return ctx.JSON(http.StatusNotFound, errors.ToErrorResponse(err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}
