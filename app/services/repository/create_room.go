package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateRoom(model model.RoomInsert, session *gorm.DB) error {
	if err := session.Table("room").Create(&model).Error; err != nil {
		return err
	}
	return nil
}
