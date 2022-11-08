package repository

import (
	"fmt"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (r mysqlRepository) SelectedPlanFixDate(req request.SelectedPlanFixDate) error {
	now := pkg.GetDatetime()
	sql := fmt.Sprintf(
		`
		UPDATE reportEngage
		SET
			selectedDate = %v, updateBy = %v, updateAt = '%v'
		WHERE
			engageId = %v
		`,
		req.SelectedDate,
		req.UpdateBy,
		now,
		req.EngageId)
	if err := r.conn.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
