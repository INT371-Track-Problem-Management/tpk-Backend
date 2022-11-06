package validator

import (
	"net/http"
	"time"
	jwt "tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (v validator) CustomerValidation(ctx echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		authentication := jwt.DecodeJWT(c)
		app := new(jwt.CheckCustomerApplication)

		if authentication.Expire < float64(time.Now().Unix()) {
			app.Id = authentication.Id
			app.Token = "Token is expired"
			app.Status = false
			res := map[string]interface{}{
				"id":      app.Id,
				"messgae": app.Token,
				"status":  app.Status,
			}
			return c.JSON(http.StatusBadRequest, res)
		}
		if authentication.Role != "C" || authentication.Status == false {
			app.Id = authentication.Id
			app.Token = "Token can't use"
			app.Status = false
			res := map[string]interface{}{
				"id":      app.Id,
				"messgae": app.Token,
				"status":  app.Status,
			}
			return c.JSON(http.StatusBadRequest, res)
		}
		app.Id = authentication.Id
		app.Token = "Token can use"
		app.Status = true

		return ctx(c)
	}
}
