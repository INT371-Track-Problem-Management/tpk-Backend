package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateReporEngage(model model.InsertReportEngage, session *gorm.DB) (*int, error) {
	var engageId int
	if err := session.Table("reportEngage").Create(&model).Error; err != nil {
		return nil, err
	}
	if err := session.Table("reportEngage").Select("engageId").Where("reportId = ?", model.ReportId).Where("createBy = ?", model.CreateBy).Where("createAt = ?", model.CreateAt).Scan(&engageId).Error; err != nil {
		return nil, err
	}
	return &engageId, nil
}
