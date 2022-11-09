package repository

import (
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (r mysqlRepository) AssignJobMaintainer(req request.AssignReport) error {
	now := pkg.GetDatetime()
	sql := `
	UPDATE reportEngage
	SET
		maintainerId = ?,
		updateAt = ?,
		updateBy = ?
	WHERE
		engageId = ?
	`
	if err := r.conn.Exec(sql, req.MaintainerId, now, req.UpdateBy, req.EngageId).Error; err != nil {
		return err
	}
	return nil
}
