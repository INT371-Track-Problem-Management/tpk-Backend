package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAssignFixReport(ctx echo.Context, conn *gorm.DB, req request.AssignReport) error {
	return service.CreateAssignFixReport(ctx, conn, req)
}
