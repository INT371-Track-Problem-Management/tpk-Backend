package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"

	"gorm.io/gorm"
)

func (r mysqlRepository) reportStatus(status request.ReportStatus, conn *gorm.DB) error {
	err := conn.Table("reportStatus").Create(status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r mysqlRepository) ReportStatusByReportId(reportId string) (*[]model.ReportStatus, error) {
	var status []model.ReportStatus
	sql := fmt.Sprintf(`
		SELECT 
			rs.statusId,
			rs.reportId,
			sm.status,
			rs.createAt
		FROM
			reportStatus rs
		LEFT JOIN
			statusMaster sm 
		ON
			rs.status = sm.statusMasterId 
		WHERE
			rs.reportId = %v
		ORDER BY
			rs.createAt DESC;
	`, reportId)
	err := r.conn.Raw(sql).Scan(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}
