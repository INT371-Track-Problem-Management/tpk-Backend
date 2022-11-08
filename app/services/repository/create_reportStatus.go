package repository

import (
	"tpk-backend/app/models/request"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateReportStatus(model request.ReportStatus, session *gorm.DB) error {
	if err := session.Table("reportStatus").Create(&model).Error; err != nil {
		return err
	}
	return nil
}
