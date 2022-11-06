package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) FixdateByEngageId(engage model.ReportEngage) (*[]model.Fixdate, error) {
	var fixdate []model.Fixdate
	sql := fmt.Sprintf(`
	SELECT
		*
	FROM
		fixDate
	WHERE
		engageId = %v
	AND
		step = %v
	`,
		engage.EngageId,
		engage.Step)
	if err := r.conn.Raw(sql).Scan(&fixdate).Error; err != nil {
		return nil, err
	}
	fmt.Println(fixdate[0])
	return &fixdate, nil
}
