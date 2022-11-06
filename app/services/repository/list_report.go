package repository

import (
	"fmt"
	"tpk-backend/app/models/request"
	"tpk-backend/app/models/response"
)

func (r mysqlRepository) ListReport(fillter *request.FillterReport) (*[]response.ReportList, error) {
	var reports []response.ReportList
	sql :=
		`
	SELECT 
		r.*,
		r2.roomNum,
		fd.step,
		fd.date
	FROM 
		reports r
	LEFT JOIN
		reportEngage re
	ON
		r.reportId = re.reportId
	LEFT JOIN 
		fixDate fd
	ON
		re.selectedDate = fd.Id
	LEFT JOIN
		room r2
	ON 
		r2.roomId = r.roomId
	WHERE
		1=1
	`
	if fillter.RoomId != "" {
		sql += fmt.Sprintf(` AND r.roomId = %v`, fillter.RoomId)
	}
	if fillter.CustomerId != "" {
		sql += fmt.Sprintf(` AND r.createBy = %v`, fillter.CustomerId)
	}
	if err := r.conn.Raw(sql).Scan(&reports).Error; err != nil {
		return nil, err
	}
	return &reports, nil
}
