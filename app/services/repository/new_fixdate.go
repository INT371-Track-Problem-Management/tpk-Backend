package repository

import (
	"fmt"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"

	"gorm.io/gorm"
)

func (r mysqlRepository) EditEngage(req request.ReportEngage, session *gorm.DB) error {
	now := pkg.GetDatetime()
	sql := fmt.Sprintf(
		`
		UPDATE reportEngage
		SET selectedDate = NULL, step = %v, updateBy = %v, updateAt = '%v'
		WHERE
			engageId = %v
		`,
		req.Step,
		req.UpdatedBy,
		now,
		req.EngageId,
	)

	if err := session.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
