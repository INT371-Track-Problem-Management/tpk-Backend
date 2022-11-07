package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) GetRoomWithCustomerId(ctx echo.Context) error {
	param := ctx.Param("customerId")
	customerId, _ := strconv.Atoi(param)
	rooms, err := c.service.GetRoomWithCustomerId(customerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, rooms)
}

func (c controllerTPK) GetAllRoomWithCustomerByBuildingId(ctx echo.Context) error {
	param := ctx.Param("buildingId")
	buildingId, _ := strconv.Atoi(param)

	rooms, err := c.service.GetAllRoomWithCustomerByBuildingId(buildingId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, rooms)
}
