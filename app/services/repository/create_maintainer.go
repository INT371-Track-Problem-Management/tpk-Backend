package repository

import (
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) CreateMaintainer(req model.AddMaintainer) error {
	if err := r.conn.Table("maintainer").Create(&req).Error; err != nil {
		return err
	}
	return nil
}
