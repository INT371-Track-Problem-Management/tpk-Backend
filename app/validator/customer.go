package validator

import (
	"net/http"
	"time"
	"tpk-backend/app/authentication"

	"github.com/labstack/echo/v4"
)

func CustomerValidation(ctx echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		jwt := authentication.DecodeJWT(c)
		app := new(authentication.CheckCustomerApplication)
		if jwt.Expire < float64(time.Now().Unix()) {
			app.Id = jwt.Id
			app.Token = "Token is expired"
			app.Status = false
			res := map[string]interface{}{
				"id":      app.Id,
				"messgae": app.Token,
				"status":  app.Status,
			}
			return c.JSON(http.StatusBadRequest, res)
		}
		if jwt.Role != "C" && jwt.Status == false {
			app.Id = jwt.Id
			app.Token = "Token can't use"
			app.Status = false
			res := map[string]interface{}{
				"id":      app.Id,
				"messgae": app.Token,
				"status":  app.Status,
			}
			return c.JSON(http.StatusBadRequest, res)
		}
		app.Id = jwt.Id
		app.Token = "Token can use"
		app.Status = true

		return ctx(c)
	}
}