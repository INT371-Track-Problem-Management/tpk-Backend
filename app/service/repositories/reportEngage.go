package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB, dormId int) (*[]entity.ReportEngage, error) {
	var data []entity.ReportEngage
	err := conn.Table("reportEngage").Where("dormId = ?", dormId).Find(&data).Error
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

func GetReportEngageByReportId(ctx echo.Context, conn *gorm.DB, reportId int) (*entity.ReportEngage, error) {
	var data entity.ReportEngage
	err := conn.Table("reportEngage").Where("reportId = ?", reportId).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ReportEngageInsert(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	stmt := conn.Begin()
	var err error
	err = stmt.Exec(`
	INSERT INTO reportEngage (date1, date2, date3, date4, reportId, dormId, updatedBy)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`, req.Date1, req.Date2, req.Date3, req.Date4, req.ReportId, req.DormId, req.UpdatedBy).Error
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	var id int
	err = stmt.Table("reportEngage").Select("engageId").Where("reportId = ?", req.ReportId).Scan(&id).Error
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	stmt.Commit()
	return &id, nil
}

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, reportId int) (*response.ReportEngageJoinReport, error) {
	result := new(response.ReportEngageJoinReport)
	sql := fmt.Sprintf(
		`
		SELECT 
			re.engageId,
			re.date1,
			re.date2,
			re.date3,
			re.date4,
			re.selectedDate,
			re.reportId,
			r.title,
			r.categoriesReport,
			r.reportDes,
			r.status,
			r.successDate,
			r.reportDate,
			r.createdBy
		FROM reportEngage re 
		JOIN reports r 
		ON re.reportId = r.reportId 
		WHERE re.reportId = %v
		`, reportId)

	err := conn.Raw(sql).Scan(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req request.SelectedPlanFixDate) error {
	stmt := conn.Begin()
	err := stmt.Exec(
		`
		UPDATE reportEngage
		SET selectedDate = ?
		WHERE engageId  = ?
		`, req.SelectedDate, req.EngageId).Error
	if err != nil {
		return err
	}
	stmt.Commit()

	return nil
}
