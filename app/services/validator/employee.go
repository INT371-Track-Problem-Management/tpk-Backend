package validator

import (
	"net/http"
	"time"
	jwt "tpk-backend/app/jwt"

	"github.com/labstack/echo/v4"
)

func (v validator) EmployeeValidation(ctx echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authentication := jwt.DecodeJWT(c)
		status := v.StatusToken(jwt.GetTokenFromHeadler(c))
		if !status {
			c.JSON(http.StatusUnauthorized, "Token is inactive")
		}

		app := new(jwt.CheckOwnerApplication)

		if authentication.Expire < float64(time.Now().Unix()) {
			app.Id = authentication.Id
			app.Token = "Token is expired"
			app.Status = false
			res := map[string]interface{}{
				"id":      app.Id,
				"messgae": app.Token,
				"status":  app.Status,
			}
			return c.JSON(http.StatusUnauthorized, res)
		}
		if authentication.Role == "E" || authentication.Role == "A" || !authentication.Status {
			app.Id = authentication.Id
			app.Id = authentication.Id
			app.Token = "Token can use"
			app.Status = true
			return ctx(c)
		}

		app.Token = "Token can't use"
		app.Status = false
		res := map[string]interface{}{
			"id":      app.Id,
			"messgae": app.Token,
			"status":  app.Status,
		}

		return c.JSON(http.StatusUnauthorized, res)
	}
}
