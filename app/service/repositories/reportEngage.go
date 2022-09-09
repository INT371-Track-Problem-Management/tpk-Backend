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

func ReportEngageInsert(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	stmt := conn.Begin()
	var err error
	err = stmt.Table("reportEngage").Create(&req).Error
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

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, customerId int) (*response.ReportEngageJoinReport, error) {
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
		WHERE r.createdBy = %v
		`, customerId)

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
