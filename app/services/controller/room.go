package controller

import (
	"net/http"
	"strconv"
	"tpk-backend/app/models/request"

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

func (c controllerTPK) CreateRoom(ctx echo.Context) error {
	req := new(request.RoomInsert)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.CreateRoom(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
