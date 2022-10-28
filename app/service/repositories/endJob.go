package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndJobReport(ctx echo.Context, conn *gorm.DB, req entity.EndJobReport) error {
	var err error
	sql1 := fmt.Sprintf(
		`
		UPDATE reports
		SET status = 'S7', updateAt = '%v', updateBy = '%v'
		WHERE reportId = %v
		`, req.DateOfIssue, req.UpdateBy, req.ReportId)
	err = conn.Exec(sql1).Error
	if err != nil {
		return err
	}

	sql2 := fmt.Sprintf(
		`
	INSERT INTO reviewReports (des, reportId, score)
	VALUES ('%v', %v, %v);
	`, req.Des, req.ReportId, req.Score)

	err = conn.Exec(sql2).Error
	if err != nil {
		return err
	}
	return nil
}
