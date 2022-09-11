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
	sql1 := fmt.Sprintf(
		`
		UPDATE reports
		SET status = 'S7'
		WHERE reportId = %v
		`, req.ReportId)
	err = stmt.Exec(sql1).Error
	if err != nil {
		return err
	}

	sql2 := fmt.Sprintf(
		`
	INSERT INTO reviewReports (des, reportId, score)
	VALUES ('%v', %v, %v);
	`, req.Des, req.ReportId, req.Score)

	err = stmt.Exec(sql2).Error
	if err != nil {
		return err
	}

	// err = stmt.Table("historyReport").Where("reportId = ?", req.ReportId).Update("dateOfIssue = ?", req.DateOfIssue).Error
	sql3 := fmt.Sprintf(
		`
		UPDATE historyReport 
		SET dateOfIssue = DATE('%v')
		WHERE reportId = %v;
	`, req.DateOfIssue, req.ReportId)
	err = stmt.Exec(sql3).Error
	if err != nil {
		return err
	}

	stmt.Commit()
	return nil
}
