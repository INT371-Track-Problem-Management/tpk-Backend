package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddEmployeeInDorm(ctx echo.Context, conn *gorm.DB, req request.AddEmpInDorm) error {
	return service.AddEmployeeInDorm(ctx, conn, req)
}
