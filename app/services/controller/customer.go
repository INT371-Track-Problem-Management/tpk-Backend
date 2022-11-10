package controller

import (
	"net/http"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchCustomerByEmail(ctx echo.Context) error {
	user := jwt.DecodeJWT(ctx)

	customer, err := c.service.FetchCustomerByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	respone := map[string]interface{}{
		"customer": customer,
	}
	return ctx.JSON(http.StatusOK, respone)
}
