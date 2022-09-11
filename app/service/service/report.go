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
		ReportDate:       data.ReportDate,
		SuccessDate:      data.SuccessDate,
		Status:           data.Status,
		CreatedBy:        data.CreatedBy,
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
		ReportDate:       timenow,
		CreatedBy:        req.CreatedBy,
	}

	reportid, err := repositories.ReportInsert(ctx, conn, data)
	if err != nil {
		return nil, err
	}

	roomProfile, err := repositories.GetRoomWithCustomerByCustomerId(ctx, conn, req.CreatedBy)
	if err != nil {
		return nil, err
	}

	history := entity.CreateHistoryReport{
		ReportId:   *reportid,
		ReportDate: timenow,
		RoomId:     roomProfile.RoomId,
		CustomerId: req.CreatedBy,
		DormId:     roomProfile.DormId,
	}

	err = repositories.CreatedHistoryReport(ctx, conn, history)
	if err != nil {
		return nil, err
	}

	// cus, err := repositories.GetCustomerById(ctx, conn, req.CreatedBy)
	// if err != nil {
	// 	return nil, err
	// }

	// if cus.Email != "" {
	// 	rps := config.LoadReportSend()
	// 	pkg.SSLemail(&cus.Email, rps.Subject, rps.Body)
	// }
	// fmt.Printf("customerId %v is not have email", req.CreatedBy)

	return reportid, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) (string, error) {
	err := repositories.ReportChangeStatus(ctx, conn, req)
	if err != nil {
		return "Can not update", err
	}
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

func ReportByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Report, error) {
	res, err := repositories.ReportByDormId(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
