package controller

import (
	"net/http"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) Login(ctx echo.Context) error {
	var err error
	user := new(request.User)
	err = ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	authen, err := c.service.Login(*user)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	response := map[string]string{
		"token": authen.Token,
		"name":  authen.Name,
	}
	return ctx.JSON(http.StatusOK, response)
}
