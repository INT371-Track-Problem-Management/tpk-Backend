package controller

import (
	"net/http"
	"strconv"
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

func (c controllerTPK) RemoveCustomerFromRoom(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.QueryParam("id"))
	if err := c.service.RemoveCustomerFromRoom(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
