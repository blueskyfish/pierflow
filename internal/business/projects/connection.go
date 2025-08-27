package projects

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) ProjectEventConnect(ctx echo.Context) error {
	err := pm.eventClient.Listen(ctx)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}

func (pm *ProjectManager) ProjectEventPing(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("user is required"))
	}
	err := pm.eventClient.SendTo(userId, "", "ping")
	if err != nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse(err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}
