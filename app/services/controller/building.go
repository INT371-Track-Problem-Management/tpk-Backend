package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) AllBuilding(ctx echo.Context) error {
	buildings, err := c.service.AllBuilding()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"buildings": buildings,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c controllerTPK) CreateBuilding(ctx echo.Context) error {
	req := new(request.BuildingInsert)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	buildingId, err := c.service.CreateBuilding(*req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]interface{}{
		"buildingId": buildingId,
	}
	return ctx.JSON(http.StatusOK, response)
}
