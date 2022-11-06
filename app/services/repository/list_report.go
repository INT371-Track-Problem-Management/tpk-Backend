package repository

import (
	"fmt"
	"tpk-backend/app/models/response"
)

func (r mysqlRepository) ListReport() (*[]response.ReportList, error) {
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
		r2.roomId = r.roomId;
	`
	if err := r.conn.Raw(sql).Scan(&reports).Error; err != nil {
		return nil, err
	}
	fmt.Println(reports)
	return &reports, nil
}
