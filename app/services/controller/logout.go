package controller

import (
	"net/http"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) LogoutToken(ctx echo.Context) error {
	token := jwt.GetTokenFromHeadler(ctx)
	if err := c.service.LogoutToken(token); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	response := map[string]string{
		"message": "success",
	}
	return ctx.JSON(http.StatusOK, response)
}
