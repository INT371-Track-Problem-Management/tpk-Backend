package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EmployeeById(ctx echo.Context, conn *gorm.DB, id int) (*entity.Employee, error) {
	emp, err := service.EmployeeById(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
