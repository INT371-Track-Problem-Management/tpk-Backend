package repositories

import (
	"log"
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

func ReportByCreatedBy(ctx echo.Context, conn *gorm.DB, req request.ReportByCreatedBy) (*[]entity.Report, error) {
	var report []entity.Report
	err := conn.Table("reports").Where(`createdBy = ?`, req.CreatedBy).Find(&report).Error
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

func ReportInsert(ctx echo.Context, conn *gorm.DB, req entity.ReportInsert) (*int, error) {
	var err error
	err = conn.Table("reports").Create(&req).Error
	if err != nil {
		return nil, err
	}
	log.Println("cratedBy")
	log.Println(req.CreatedBy)
	var id int
	err = conn.Table("reports").Select("reportId").Where("title = ?", req.Title).Where("createdBy = ?", req.CreatedBy).Where("reportDate = ?", req.ReportDate).Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) error {
	err := conn.Exec("UPDATE reports SET status = ? WHERE reportId = ?", req.Status, req.ReportId).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteReportById(ctx echo.Context, conn *gorm.DB, req request.Report) error {
	var err error
	session := conn.Begin()
	err = session.Exec("DELETE FROM reviewReports WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM reportEngage WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM assignReport WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	err = session.Exec("DELETE FROM reports WHERE reportId = ?", req.ReportId).Error
	if err != nil {
		return err
	}
	session.Commit()
	return nil
}

func ReportByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Report, error) {
	var data *[]entity.Report
	err := conn.Table("report").Where("dormId = ?", dormId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
