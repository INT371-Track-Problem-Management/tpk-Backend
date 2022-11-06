package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) SelectedPlanFixDate(ctx echo.Context) error {
	req := new(request.SelectedPlanFixDate)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.SelectedPlanFixDate(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
