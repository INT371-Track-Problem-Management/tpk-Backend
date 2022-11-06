package repository

import (
	"fmt"
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) FixdateByEngageId(engageId int) (*[]model.Fixdate, error) {
	var fixdate []model.Fixdate
	sql := fmt.Sprintf(`
	SELECT
		*
	FROM
		fixDate
	WHERE
		engageId = %v
	`, engageId)
	if err := r.conn.Raw(sql).Scan(&fixdate).Error; err != nil {
		return nil, err
	}
	fmt.Println(fixdate[0])
	return &fixdate, nil
}
