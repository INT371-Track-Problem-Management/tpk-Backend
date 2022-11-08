package repository

import (
	"tpk-backend/app/models/model"

	"gorm.io/gorm"
)

func (r mysqlRepository) CreateReport(model model.ReportInsert, session *gorm.DB) (*int, error) {
	var reportId int
	if err := session.Table("reports").Create(&model).Error; err != nil {
		return nil, err
	}
	if err := session.Table("reports").Select("reportId").Where("title = ?", model.Title).Where("createBy = ?", model.CreateBy).Where("createAt = ?", model.CreateAt).Scan(&reportId).Error; err != nil {
		return nil, err
	}
	return &reportId, nil
}
