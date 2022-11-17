package repository

import (
	"tpk-backend/app/models/model"
)

func (r mysqlRepository) CreateProfileMedia(req model.ProfileMedia) error {
	if err := r.conn.Table("profileMedia").Create(&req).Error; err != nil {
		return err
	}
	return nil
}
