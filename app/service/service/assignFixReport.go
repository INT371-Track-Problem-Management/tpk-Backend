package service

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAssignFixReport(ctx echo.Context, conn *gorm.DB, req request.AssignReport) error {
	timenow := pkg.GetDatetime()
	model := entity.AssignReport{
		EngageId:     req.EngageId,
		ReportId:     req.ReportId,
		MaintainerId: req.MaintainerId,
		UpdateBy:     req.UpdateBy,
		UpdateAt:     timenow,
	}
	return repositories.CreateAssignFixReport(ctx, conn, model)
}
