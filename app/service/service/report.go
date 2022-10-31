package service

import (
	"errors"
	"fmt"
	"tpk-backend/app/constants"
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/pkg"
	"tpk-backend/app/pkg/config"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var fileConfig = config.LoadFileConfig()

func Report(ctx echo.Context, conn *gorm.DB) (*[]entity.ReportJoinEngage, error) {
	report, err := repositories.Report(ctx, conn)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func ReportById(ctx echo.Context, conn *gorm.DB, req request.Report) (*response.Report, error) {
	data, err := repositories.ReportById(ctx, conn, req.ReportId)
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
		RoomId:           data.RoomId,
		RoomNum:          data.RoomNum,
		BuildingId:       data.BuildingId,
		SelectedDate:     data.SelectedDate,
	}
	return res, nil
}

func ReportByCreatedBy(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.Report, error) {
	res, err := repositories.ReportByCreatedBy(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ReportInsert(ctx echo.Context, conn *gorm.DB, req request.ReportInsert, filename string) (*int, error) {
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
		RoomId:           req.RoomId,
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

	media := entity.ReportMediaInsert{
		ReportId: *reportid,
		Name:     filename,
		Url:      fileConfig.URL + "/" + fileConfig.Bucket + "/" + filename,
		CreateAt: timenow,
		UpdateAt: timenow,
	}

	err = repositories.UploadFile(ctx, session, media)
	if err != nil {
		return nil, err
	}

	session.Commit()

	customer, err := repositories.GetCustomerById(ctx, conn, req.CreateBy)
	if err != nil {
		return nil, err
	}

	err = pkg.Smtp2(constants.SUBJECT_EMAIL_SENDING_REPORT, customer.Email, "ส่งการรายงาน")
	if err != nil {
		return nil, err
	}

	return reportid, nil
}

func ReportChangeStatus(ctx echo.Context, conn *gorm.DB, req request.ReportChangeStatus) error {
	timenow := pkg.GetDatetime()
	status := request.ReportStatus{
		ReportId:  req.ReportId,
		Status:    req.Status,
		UpdateAt:  timenow,
		UpdateBy:  req.UpdateBy,
		CreatedAt: timenow,
	}
	insert := entity.ReportChangeStatus{
		ReportId: req.ReportId,
		Status:   req.Status,
		UpdateAt: timenow,
		UpdateBy: req.UpdateBy,
	}
	session := conn.Begin()
	err := repositories.ReportChangeStatus(ctx, session, insert)
	if err != nil {
		return err
	}
	err = repositories.ReportStatus(ctx, session, status)
	if err != nil {
		return err
	}
	session.Commit()

	report, err := repositories.ReportById(ctx, conn, req.ReportId)
	if err != nil {
		return err
	}

	email := repositories.GetemailByCustomerId(ctx, conn, report.CreateBy)
	err = pkg.UpdateStatus(*email, report.ReportId, report.Title, report.Status)
	if err != nil {
		return err
	}

	return nil
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

func ReportByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*[]entity.ReportJoinEngage, error) {
	reports, err := repositories.ReportByRoomId(ctx, conn, roomId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func ReportStatusByReportId(ctx echo.Context, conn *gorm.DB, reportId string) (*[]response.ReportStatus, error) {
	var list []response.ReportStatus
	status, err := repositories.ReportStatusByReportId(ctx, conn, reportId)
	if err != nil {
		return nil, err
	}
	if len(*status) == 0 {
		return nil, errors.New("status not found")
	}
	for _, v := range *status {
		report := response.ReportStatus{
			StatusId:  v.StatusId,
			ReportId:  v.ReportId,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
		}
		list = append(list, report)
	}
	return &list, nil
}

func ReportListForCustomer(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.ReportJoinEngage, error) {
	reports, err := repositories.ReportListForCustomer(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}
