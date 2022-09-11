package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndJobReport(ctx echo.Context, conn *gorm.DB, req entity.EndJobReport) error {
	var err error
	stmt := conn.Begin()

	err = stmt.Table("reports").Where("reportId = ?", req.ReportId).Update("status = ?", "S7").Error
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(
		`
	INSERT INTO reviewReports (des, reportId, score)
	VALUES (%v, %v, %v);
	`, req.Des, req.ReportId, req.Score)

	err = stmt.Exec(sql).Error
	if err != nil {
		return err
	}

	err = stmt.Table("historyReport").Where("reportId = ?", req.ReportId).Update("dateOfIssue = ?", req.DateOfIssue).Error
	if err != nil {
		return err
	}

	stmt.Commit()
	return nil
}
