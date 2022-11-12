package repository

import (
	"fmt"
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) EndJobReport(session *gorm.DB, req model.EndJobReport) error {
	sql1 := fmt.Sprintf(
		`
		UPDATE reports
		SET status = 'S7', updateAt = '%v', updateBy = '%v'
		WHERE reportId = %v
		`, req.DateOfIssue, req.UpdateBy, req.ReportId)
	if err := session.Exec(sql1).Error; err != nil {
		return err
	}

	sql2 := fmt.Sprintf(
		`
	INSERT INTO reviewReports (des, reportId, score)
	VALUES ('%v', %v, %v);
	`, req.Des, req.ReportId, req.Score)
	if err := session.Exec(sql2).Error; err != nil {
		return err
	}
	return nil
}
