package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetHistoryByCustomerId(ctx echo.Context, conn *gorm.DB, customerId int64) (*[]entity.HistoryReport, error) {
	id := int(customerId)
	return service.GetHistoryByCustomerId(ctx, conn, id)
}

func GetHistoryByEmployeeId(ctx echo.Context, conn *gorm.DB, employeeId int64) (*[]entity.HistoryReport, error) {
	id := int(employeeId)
	return service.GetHistoryByEmployeeId(ctx, conn, id)
}

func GetHistoryByHistoryId(ctx echo.Context, conn *gorm.DB, historyId int64) (*entity.HistoryReport, error) {
	id := int(historyId)
	return service.GetHistoryByHistoryId(ctx, conn, id)
}
