package controller

import (
	"net/http"
	"tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (c controllerTPK) FetchProfile(ctx echo.Context) error {
	user := jwt.DecodeJWT(ctx)

	if user.Role == "E" || user.Role == "A" {
		employee, err := c.service.FetchEmployeeByEmail(user.Email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		respone := map[string]interface{}{
			"employee": employee,
		}
		return ctx.JSON(http.StatusOK, respone)
	}

	if user.Role == "C" {
		customer, err := c.service.FetchCustomerByEmail(user.Email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		respone := map[string]interface{}{
			"customer": customer,
		}
		return ctx.JSON(http.StatusOK, respone)
	}
	return ctx.JSON(http.StatusNoContent, "No content")
}

func (c controllerTPK) UploadProfilePic(ctx echo.Context) error {
	email := ctx.Param("email")
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "need file image")
	}
	if _, err := c.service.GetUser(email); err != nil {
		ctx.JSON(http.StatusBadRequest, "email not found")
	}
	if err := c.service.CreateProfileMedia(file, email); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, "success")
}
