package service

import (
	"fmt"
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Report(ctx echo.Context, conn *gorm.DB) (*[]entity.Report, error) {
	report, err := repositories.Report(ctx, conn)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func ReportById(ctx echo.Context, conn *gorm.DB, req request.Report) (*response.Report, error) {
	data, err := repositories.ReportById(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	res := &response.Report{
		ReportId:         data.ReportId,
		Title:            data.Title,
		CategoriesReport: data.CategoriesReport,
		ReportDes:        data.ReportDes,
		Status:           data.Status,
		UpdateAt:         data.UpdateAt,
		UpdateBy:         data.UpdateBy,
		CreateAt:         data.CreateAt,
		CreateBy:         data.CreateBy,
	}
	return res, nil
}

func ReportByCreatedBy(ctx echo.Context, conn *gorm.DB, req request.ReportByCreatedBy) (*[]entity.Report, error) {
	res, err := repositories.ReportByCreatedBy(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req request.ReportInsert) (*int, error) {
	timenow := pkg.GetDatetime()
	data := entity.ReportInsert{
		Title:            req.Title,
		CategoriesReport: req.CategoriesReport,
		ReportDes:        req.ReportDes,
		Status:           req.Status,
		CreateAt:         timenow,
		CreateBy:         req.CreateBy,
		UpdateAt:         timenow,
		UpdateBy:         req.CreateBy,
	}

	session := conn.Begin()
	reportid, err := repositories.ReportInsert(ctx, session, data)
	if err != nil {
		return nil, err
	}

	status := request.ReportStatus{
		ReportId:  *reportid,
		Status:    req.Status,
		UpdateAt:  timenow,
		UpdateBy:  req.CreateBy,
		CreatedAt: timenow,
	}

	err = repositories.ReportStatus(ctx, session, status)
	if err != nil {
		return nil, err
	}

	session.Commit()

	return reportid, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) (string, error) {
	timenow := pkg.GetDatetime()
	status := request.ReportStatus{
		ReportId:  req.ReportId,
		Status:    req.Status,
		UpdateAt:  timenow,
		UpdateBy:  req.EmployeeId,
		CreatedAt: timenow,
	}
	insert := entity.ReportChangeStatus{
		ReportId:   req.ReportId,
		Status:     req.Status,
		UpdateAt:   timenow,
		EmployeeId: req.EmployeeId,
	}
	session := conn.Begin()
	err := repositories.ReportChangeStatus(ctx, session, insert)
	if err != nil {
		return "Can not update", err
	}
	err = repositories.ReportStatus(ctx, session, status)
	if err != nil {
		return "Can not update", err
	}
	session.Commit()
	return "Update success", nil
}

func DeleteReportById(ctx echo.Context, conn *gorm.DB, req request.Report) error {
	err := repositories.DeleteReportById(ctx, conn, req)
	if err != nil {
		return err
	}
	fmt.Printf("Delete report id %v success", req.ReportId)
	return nil
}

func YearConfig(ctx echo.Context, conn *gorm.DB) (*response.Year, error) {
	year, err := repositories.YearConfig(ctx, conn)
	if err != nil {
		return nil, err
	}
	return year, nil
}

func ReportByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*[]entity.Report, error) {
	reports, err := repositories.ReportByRoomId(ctx, conn, roomId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}
