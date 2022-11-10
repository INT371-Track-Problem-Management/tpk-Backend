package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) ChangePassword(ctx echo.Context) error {
	req := new(request.ChangePassword)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if err := c.service.ChangePassword(*req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "success")
}
