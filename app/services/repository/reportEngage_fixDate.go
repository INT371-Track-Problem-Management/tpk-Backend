package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) ReportEnagegeByReportId(reportId string) (*model.ReportEngage, error) {
	var engage model.ReportEngage
	sql := fmt.Sprintf(`
	SELECT
		*
	FROM
		reportEngage
	WHERE
		reportId = %v
	`, reportId)
	if err := r.conn.Raw(sql).Scan(&engage).Error; err != nil {
		return nil, err
	}
	return &engage, nil
}
