package repository

import (
	"fmt"
	"tpk-backend/app/models/response"
)

func (r mysqlRepository) ReportDetailById(reportId int) (*response.ReportDetailById, error) {
	var report response.ReportDetailById
	sql := fmt.Sprintf(
		`
		SELECT
			reports.*,
			room.roomNum,
			reportMedia.id as imageId
		FROM
			reports 
		LEFT JOIN
			reportMedia
		ON
			reports.reportId = reportMedia.reportId
		LEFT JOIN
			room
		ON
			reports.roomId = room.roomId  
		WHERE
			reports.reportId = %v;
		`,
		reportId,
	)
	if err := r.conn.Raw(sql).Scan(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}
