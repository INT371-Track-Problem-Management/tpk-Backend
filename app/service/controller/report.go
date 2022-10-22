package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Report(ctx echo.Context, conn *gorm.DB) (*[]entity.Report, error) {
	res, err := service.Report(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetReportByCreatedBy(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.Report, error) {
	res, err := service.ReportByCreatedBy(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportById(ctx echo.Context, conn *gorm.DB, req request.Report) (*response.Report, error) {
	res, err := service.ReportById(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req request.ReportInsert) (*int, error) {
	res, err := service.ReportInsert(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) (*string, error) {
	res, err := service.ReportChangeStatus(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DeleteReportById(ctx echo.Context, conn *gorm.DB, req request.Report) error {
	err := service.DeleteReportById(ctx, conn, req)
	if err != nil {
		return err
	}
	return nil
}

func YearConfig(ctx echo.Context, conn *gorm.DB) (*response.Year, error) {
	year, err := service.YearConfig(ctx, conn)
	if err != nil {
		return nil, err
	}
	return year, nil
}

func ReportByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*[]entity.ReportJoinEngage, error) {
	reports, err := service.ReportByRoomId(ctx, conn, roomId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func ReportStatusByReportId(ctx echo.Context, conn *gorm.DB, reportId string) (*[]response.ReportStatus, error) {
	status, err := service.ReportStatusByReportId(ctx, conn, reportId)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func ReportListForCustomer(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.ReportJoinEngage, error) {
	reports, err := service.ReportListForCustomer(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}
