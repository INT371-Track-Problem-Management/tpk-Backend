package repository

import (
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) ReportDetailById(reportId string) (*model.Report, error) {
	var report model.Report
	if err := r.conn.Table("reports").Where("reportId = ?", reportId).Find(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}
