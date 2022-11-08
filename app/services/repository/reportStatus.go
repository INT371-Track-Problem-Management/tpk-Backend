package repository

import (
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error) {
	var status []model.ReportStatus
	if err := r.conn.Table("reportStatus").Order("createAt DESC").Where("reportId = ?", reportId).Scan(&status).Error; err != nil {
		return nil, err
	}
	return &status, nil
}
