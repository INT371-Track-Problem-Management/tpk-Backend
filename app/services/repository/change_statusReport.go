package repository

import (
	"tpk-backend/app/models/request"

	"gorm.io/gorm"
)

func (r mysqlRepository) ChangeStatusReport(req request.ReportStatus, session *gorm.DB) error {
	if err := r.conn.Exec("UPDATE reports SET status = ?, updateBy = ?, updateAt = ? WHERE reportId = ?", req.Status, req.UpdateBy, req.UpdateAt, req.ReportId).Error; err != nil {
		return err
	}
	return nil
}
