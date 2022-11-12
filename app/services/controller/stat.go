package controller

import (
	"net/http"
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
