package controller

import (
	"net/http"
	"tpk-backend/app/jwt"
	"tpk-backend/app/models/request"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) ChangeEmail(ctx echo.Context) error {
	var err error
	req := new(request.ChangeEmail)
	err = ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	token := jwt.DecodeJWT(ctx)

	err = c.service.ChangeEmail(*req, token.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadGateway, err)
	}
	return ctx.JSON(http.StatusCreated, nil)
}
