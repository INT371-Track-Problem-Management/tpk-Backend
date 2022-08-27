package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndJobReport(ctx echo.Context, conn *gorm.DB, req request.EndJobReport) error {
	return service.EndJobReport(ctx, conn, req)
}
