package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateReportMedia(req model.ReportMedia, session *gorm.DB) error {
	if err := r.conn.Table("reportMedia").Create(&req).Error; err != nil {
		return err
	}
	return nil
}
