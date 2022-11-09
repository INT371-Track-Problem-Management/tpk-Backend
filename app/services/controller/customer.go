package controller

import (
	"net/http"
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/request"

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

func (c controllerTPK) CustomerEditProfile(ctx echo.Context) error {
	req := new(request.CustomerEditProfile)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	email := ctx.QueryParam("email")
	err = c.service.CustomerEditProfile(*req, email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "success")
}
