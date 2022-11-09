package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) AddCustomerIntoRoom(ctx echo.Context) error {
	req := new(request.RoomAddCustomer)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.AddCustomerIntoRoom(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
