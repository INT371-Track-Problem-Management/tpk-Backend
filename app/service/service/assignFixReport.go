package service

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAssignFixReport(ctx echo.Context, conn *gorm.DB, req request.AssignReport) error {
	return repositories.CreateAssignFixReport(ctx, conn, req)
}
