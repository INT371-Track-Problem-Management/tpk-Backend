package service

import (
	"errors"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddEmployeeInBuilding(ctx echo.Context, conn *gorm.DB, req request.AddEmpInBuilding) error {
	employee, err := repositories.EmployeeById(ctx, conn, req.EmployeeId)
	if err != nil {
		return err
	}

	if employee.Position != "owner" {
		return errors.New("employee_is_not_owner_position")
	}

	return repositories.AddEmployeeInBuilding(ctx, conn, req)
}
