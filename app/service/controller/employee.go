package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EmployeeById(ctx echo.Context, conn *gorm.DB, id int64) (*entity.Employee, error) {
	empId := int(id)
	emp, err := service.EmployeeById(ctx, conn, empId)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
