package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ListEmployee(ctx echo.Context, conn *gorm.DB) (*[]entity.Employee, error) {
	employees, err := repositories.ListEmployee(ctx, conn)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func EmployeeById(ctx echo.Context, conn *gorm.DB, id int) (*entity.Employee, error) {
	emp, err := repositories.EmployeeById(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
