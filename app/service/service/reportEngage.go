package service

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportEngageAll(ctx echo.Context, conn *gorm.DB) (*response.ReportEngageAll, error) {
	data, err := repositories.GetReportEngageAll(ctx, conn)
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
		EngageId:   data.EngageId,
		SelectDate: data.SelectedDate,
		Date1:      data.Date1,
		Date2:      data.Date2,
		Date3:      data.Date3,
		Date4:      data.Date4,
		ReportId:   data.ReportId,
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
