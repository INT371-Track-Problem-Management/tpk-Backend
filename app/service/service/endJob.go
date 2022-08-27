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
	}
	return repositories.EndJobReport(ctx, conn, entity)
}
