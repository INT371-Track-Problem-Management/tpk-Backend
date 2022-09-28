package service

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
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
		Date1:        data.Date1,
		Date2:        data.Date2,
		Date3:        data.Date3,
		Date4:        data.Date4,
		SelectedDate: data.SelectedDate,
		ReportId:     data.ReportId,
		BuildingId:   data.BuildingId,
		CreateBy:     data.CreateBy,
		CreateAt:     data.CreateAt,
		UpdateAt:     data.UpdateAt,
		UpdateBy:     data.UpdateBy,
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
		Date1:        data.Date1,
		Date2:        data.Date2,
		Date3:        data.Date3,
		Date4:        data.Date4,
		SelectedDate: data.SelectedDate,
		ReportId:     data.ReportId,
		BuildingId:   data.BuildingId,
		CreateBy:     data.CreateBy,
		CreateAt:     data.CreateAt,
		UpdateAt:     data.UpdateAt,
		UpdateBy:     data.UpdateBy,
	}
	return &res, nil
}

func InsertReportEngage(ctx echo.Context, conn *gorm.DB, req request.ReportEngage) (*int, error) {
	timenow := pkg.GetDatetime()
	model := entity.InsertReportEngage{
		Date1:      req.Date1,
		Date2:      req.Date2,
		Date3:      req.Date3,
		Date4:      req.Date4,
		ReportId:   req.ReportId,
		BuildingId: req.BuildingId,
		CreateBy:   req.UpdatedBy,
		CreateAt:   timenow,
		UpdateAt:   timenow,
		UpdateBy:   req.UpdatedBy,
	}
	data, err := repositories.ReportEngageInsert(ctx, conn, model)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReportEngageJoinReport(ctx echo.Context, conn *gorm.DB, reportId int) (*response.ReportEngageJoinReport, error) {
	data, err := repositories.ReportEngageJoinReport(ctx, conn, reportId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SelectedDatePlanFix(ctx echo.Context, conn *gorm.DB, req request.SelectedPlanFixDate) error {
	timenow := pkg.GetDatetime()
	model := entity.SelectedPlanFixDate{
		EngageId:     req.EngageId,
		SelectedDate: req.SelectedDate,
		UpdateBy:     req.UpdateBy,
		UpdateAt:     timenow,
	}
	err := repositories.SelectedDatePlanFix(ctx, conn, model)
	if err != nil {
		return err
	}
	return nil
}
