package service

import (
	"fmt"
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
	"tpk-backend/app/pkg/config"
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
		ReportDate:       data.ReportDate,
		SuccessDate:      data.SuccessDate,
		Status:           data.Status,
		CreatedBy:        data.CreatedBy,
	}
	return res, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req request.ReportInsert) (string, error) {
	timenow := pkg.GetDatetime()
	data := entity.ReportInsert{
		Title:            req.Title,
		CategoriesReport: req.CategoriesReport,
		ReportDes:        req.ReportDes,
		Status:           req.Status,
		ReportDate:       timenow,
		CreatedBy:        req.CreatedBy,
	}
	err := repositories.ReportInsert(ctx, conn, data)
	if err != nil {
		return "Can not insert", err
	}

	cus, err := repositories.GetUserByCustomerId(ctx, conn, req.CreatedBy)
	if err != nil {
		return "Can not find profile", err
	}

	if cus.Email != "" {
		rps := config.LoadReportSend()
		pkg.SSLemail(&cus.Email, rps.Subject, rps.Body)
	}
	fmt.Printf("customerId %v is not have email", req.CreatedBy)

	return "Insert success", nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) (string, error) {
	err := repositories.ReportChangeStatus(ctx, conn, req)
	if err != nil {
		return "Can not update", err
	}
	return "Update success", nil
}
