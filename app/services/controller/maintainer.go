package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) CreateMaintainer(ctx echo.Context) error {
	req := new(request.Maintainer)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.CreateMaintainer(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
