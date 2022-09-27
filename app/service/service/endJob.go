package service

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndJobReport(ctx echo.Context, conn *gorm.DB, req request.EndJobReport) error {
	now := pkg.GetDatetime()
	entity := entity.EndJobReport{
		Des:         req.Des,
		ReportId:    req.ReportId,
		Score:       req.Score,
		DateOfIssue: now,
		UpdateBy:    req.UpdateBy,
	}
	status := request.ReportStatus{
		ReportId:  req.ReportId,
		Status:    "S7",
		UpdateAt:  now,
		UpdateBy:  req.UpdateBy,
		CreatedAt: now,
	}

	session := conn.Begin()
	err := repositories.EndJobReport(ctx, session, entity)
	if err != nil {
		return err
	}
	err = repositories.ReportStatus(ctx, session, status)
	if err != nil {
		return err
	}
	session.Commit()
	return nil
}
