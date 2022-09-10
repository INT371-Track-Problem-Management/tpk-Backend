package controller

import (
	"errors"
	"tpk-backend/app/authentication"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB, dormId int64) (*response.ReportEngageAll, error) {
	id := int(dormId)

	jwt := authentication.DecodeJWT(ctx)

	checkDormId, err := repositories.SelectDormIdByEmployeeId(ctx, conn, jwt.Id)
	if err != nil {
		return nil, err
	}

	if *checkDormId != id {
		err := errors.New("Invalid_Token")
		return nil, err
	}

	res, err := getReportEngageAll(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetReportEngageByReportId(ctx echo.Context, conn *gorm.DB, id int64) (*response.ReportEngage, error) {
	reportId := int(id)
	return service.GetReportEngageByReportId(ctx, conn, reportId)
}

func getReportEngageAll(ctx echo.Context, conn *gorm.DB, dormId int) (*response.ReportEngageAll, error) {
	res, err := service.GetReportEngageAll(ctx, conn, dormId)
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

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, reportId int64) (*response.ReportEngageJoinReport, error) {
	id := int(reportId)
	data, err := service.ReportEngageJoinReport(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req request.SelectedPlanFixDate) error {
	return service.SelectedDatePlanFix(ctx, conn, req)
}
