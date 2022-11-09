package controller

import (
	"net/http"
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) EditProfile(ctx echo.Context) error {
	var role string
	user := jwt.DecodeJWT(ctx)
	if user.Role == "C" {
		role = "customer"
	}
	if user.Role == "E" || user.Role == "A" {
		role = "employee"
	}
	req := new(request.EditProfile)
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	email := ctx.QueryParam("email")
	err = c.service.EditProfile(*req, email, role)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "success")
}
