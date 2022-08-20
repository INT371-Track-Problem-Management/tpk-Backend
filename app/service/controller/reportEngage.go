package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB) (*response.ReportEngageAll, error) {
	res, err := service.GetReportEngageAll(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetReportEngageById(ctx echo.Context, conn *gorm.DB, req request.ReportEngageById) (*response.ReportEngage, error) {
	res, err := service.GetReportEngageById(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func InsertReportEngage(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	res, err := service.InsertReportEngage(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, customerId int64) (*response.ReportEngageJoinReport, error) {
	id := int(customerId)
	return service.ReportEngageJoinReport(ctx, conn, id)
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req request.SelectedPlanFixDate) error {
	return service.SelectedDatePlanFix(ctx, conn, req)
}
