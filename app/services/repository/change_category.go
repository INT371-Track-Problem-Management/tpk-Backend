package repository

import (
	"fmt"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (r mysqlRepository) ChangeCategory(req request.ReportChangeCategory) error {
	now := pkg.GetDatetime()
	sql := fmt.Sprintf(
		`
		UPDATE reports
		SET
			categoriesReport = '%v', updateBy = %v, updateAt = '%v'
		WHERE
			reportId = %v
		`,
		req.CategoriesReport,
		req.UpdateBy,
		now,
		req.ReportId)
	if err := r.conn.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
