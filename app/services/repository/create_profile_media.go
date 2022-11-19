package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateProfileMedia(req model.ProfileMedia, session *gorm.DB) error {
	if err := session.Table("profileMedia").Create(&req).Error; err != nil {
		return err
	}
	return nil
}
