package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAssignFixReport(ctx echo.Context, conn *gorm.DB, model entity.AssignReport) error {
	stmt := conn.Begin()
	err := stmt.Exec(
		`
		UPDATE reportEngage
		SET maintainerId = ?, updateBy = ?, updateAt = ?
		WHERE reportId = ?
		`, model.MaintainerId, model.UpdateBy, model.UpdateAt, model.ReportId).Error
	if err != nil {
		return err
	}
	stmt.Commit()
	return nil
}
