package repositories

import (
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAssignFixReport(ctx echo.Context, conn *gorm.DB, req request.AssignReport) error {
	stmt := conn.Begin()
	err := stmt.Table("assignReport").Create(&req).Error
	if err != nil {
		return err
	}
	stmt.Commit()
	return nil
}
