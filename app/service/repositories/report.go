package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Report(ctx echo.Context, conn *gorm.DB) (*[]entity.Report, error) {
	var report []entity.Report
	err := conn.Table("reports").Find(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportById(ctx echo.Context, conn *gorm.DB, req request.Report) (*entity.Report, error) {
	var report entity.Report
	err := conn.Table("reports").Where("reportId = ?", req.ReportId).Find(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req entity.Report) error {
	err := conn.Table("reports").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) error {
	err := conn.Exec("UPDATE reports SET status = ? WHERE reportId = ?", req.Status, req.ReportId).Error
	if err != nil {
		return err
	}
	return nil
}
