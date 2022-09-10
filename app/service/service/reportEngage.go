package service

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB, dormId int) (*response.ReportEngageAll, error) {
	data, err := repositories.GetReportEngageAll(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	res := response.ReportEngageAll{
		Data: *data,
	}
	return &res, nil
}

func GetReportEngageById(ctx echo.Context, conn *gorm.DB, req request.ReportEngageById) (*response.ReportEngage, error) {
	data, err := repositories.GetReportEngageById(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	res := response.ReportEngage{
		EngageId:     data.EngageId,
		SelectedDate: data.SelectedDate,
		Date1:        data.Date1,
		Date2:        data.Date2,
		Date3:        data.Date3,
		Date4:        data.Date4,
		ReportId:     data.ReportId,
		DormId:       data.DormId,
		UpdatedBy:    data.UpdatedBy,
	}
	return &res, nil
}

func GetReportEngageByReportId(ctx echo.Context, conn *gorm.DB, reportId int) (*response.ReportEngage, error) {
	data, err := repositories.GetReportEngageByReportId(ctx, conn, reportId)
	if err != nil {
		return nil, err
	}
	res := response.ReportEngage{
		EngageId:     data.EngageId,
		SelectedDate: data.SelectedDate,
		Date1:        data.Date1,
		Date2:        data.Date2,
		Date3:        data.Date3,
		Date4:        data.Date4,
		ReportId:     data.ReportId,
		DormId:       data.DormId,
		UpdatedBy:    data.UpdatedBy,
	}
	return &res, nil
}

func InsertReportEngage(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	data, err := repositories.ReportEngageInsert(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, customerId int) (*response.ReportEngageJoinReport, error) {
	return repositories.ReportEngageJoinReport(ctx, conn, customerId)
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req request.SelectedPlanFixDate) error {
	return repositories.SelectedDatePlanFix(ctx, conn, req)
}
