package controller

import (
	"net/http"
	"strconv"
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

func (c controllerTPK) Maintainerlist(ctx echo.Context) error {
	maintainers, err := c.service.Maintainerlist()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, maintainers)
}
func (c controllerTPK) MaintainerById(ctx echo.Context) error {
	MaintainerId, _ := strconv.Atoi(ctx.Param("maintainerId"))
	maintainers, err := c.service.MaintainerById(MaintainerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, maintainers)
}
