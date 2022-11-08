package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateFixdate(model model.CreateFixdate, session *gorm.DB) error {
	if err := session.Table("fixDate").Create(&model).Error; err != nil {
		return err
	}
	return nil
}
