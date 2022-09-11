package pkg

import (
	"tpk-backend/app/authentication"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetUserDetail(ctx echo.Context) interface{} {
	jwt := authentication.DecodeJWT(ctx)
	if jwt.Role == "C" {
		customer, err := repositories.GetCustomerById(ctx, db, jwt.Id)
		if err != nil {
			return err
		}
		return customer
	}
	if jwt.Role == "E" {
		employee, err := repositories.EmployeeById(ctx, db, jwt.Id)
		if err != nil {
			return err
		}
		return employee
	}
	return nil
}
