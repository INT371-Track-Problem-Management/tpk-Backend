package controller

import (
	"net/http"
	"strconv"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c *controllerTPK) FetcStatDashBoard(ctx echo.Context) error {
	req := new(request.Stat)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := c.service.FetcStatDashBoard(*req)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *controllerTPK) FetchStatMaintain(ctx echo.Context) error {
	res, err := c.service.FetchStatMaintain()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *controllerTPK) FetchOverviewMaintain(ctx echo.Context) error {
	maintainerId, _ := strconv.Atoi(ctx.Param("maintainerId"))
	res, err := c.service.FetchOverviewMaintain(maintainerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
