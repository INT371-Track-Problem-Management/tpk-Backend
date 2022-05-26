package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB) (*[]entity.ReportEngage, error) {
	var data []entity.ReportEngage
	err := conn.Table("reportEngage").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetReportEngageById(ctx echo.Context, conn *gorm.DB, req request.ReportEngageById) (*entity.ReportEngage, error) {
	var data entity.ReportEngage
	err := conn.Table("reportEngage").Where("engageId = ?", req.EngageId).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ReportEngageInsert(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	var err error
	err = conn.Table("reportEngage").Create(&req).Error
	if err != nil {
		return nil, err
	}
	var id int
	err = conn.Table("reportEngage").Select("engageId").Where("reportId = ?", req.ReportId).Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}
