package controller

import (
	"net/http"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchEmployeeByEmail(ctx echo.Context) error {
	user := jwt.DecodeJWT(ctx)

	employee, err := c.service.FetchEmployeeByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	respone := map[string]interface{}{
		"employee": employee,
	}
	return ctx.JSON(http.StatusOK, respone)
}
