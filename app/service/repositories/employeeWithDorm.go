package repositories

import (
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddEmployeeInBuilding(ctx echo.Context, conn *gorm.DB, req request.AddEmpInBuilding) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Exec(`
	INSERT INTO employeeWithBuilding (employeeId, dormId)
	VALUES (?, ?)
	`,
		req.EmployeeId, req.BuildingId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	err = stmt.Exec("UPDATE employee SET position = ? WHERE employeeId = ?", "staff", req.EmployeeId).Error
	if err != nil {
		stmt.Rollback()
		return err
	}

	stmt.Commit()
	return nil
}
