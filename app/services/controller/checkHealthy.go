package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) CheckHealthy(ctx echo.Context) error {
	health, err := c.service.CheckHealthy()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &health)
}
