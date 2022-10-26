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

func ReportEngageInsert(ctx echo.Context, conn *gorm.DB, model entity.InsertReportEngage) (*int, error) {
	stmt := conn.Begin()
	var err error
	err = stmt.Table("reportEngage").Create(model).Error
	if err != nil {
		stmt.Rollback()
		return nil, err
	}
	var id int
	err = stmt.Table("reportEngage").Select("engageId").Where("reportId = ?", model.ReportId).Scan(&id).Error
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
			re.engageId as engageId,
			re.date1 as date1,
			re.date2 as date2,
			re.date3 as date3,
			re.date4 as date4,
			re.selectedDate as selectedDate,
			re.reportId as reportId,
			re.maintainerId as maintainerId,
			r.title as title,
			r.categoriesReport as categoriesReport,
			r.reportDes as reportDes,
			r.status as status,
			ro.roomNum as roomNum,
			ro.buildingId as buildingId,
			r.createAt as createAt,
			r.updateAt as updateAt,
			r.updateBy as updateBy,
			r.createBy as createBy
		FROM reportEngage re 
		JOIN reports r 
		ON re.reportId = r.reportId 
		JOIN room ro
		ON r.roomId = ro.roomId
		WHERE re.reportId = %v
		`, reportId)

	err := conn.Raw(sql).Scan(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req entity.SelectedPlanFixDate) error {
	stmt := conn.Begin()
	err := stmt.Exec(
		`
		UPDATE reportEngage
		SET selectedDate = ?, updateBy = ?, updateAt = ?
		WHERE engageId  = ?
		`, req.SelectedDate, req.UpdateBy, req.UpdateAt, req.EngageId).Error
	if err != nil {
		return err
	}
	stmt.Commit()

	return nil
}
