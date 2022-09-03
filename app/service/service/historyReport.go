package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetHistoryByCustomerId(ctx echo.Context, conn *gorm.DB, customerId int) (*[]entity.HistoryReport, error) {
	return repositories.GetHistoryByCustomerId(ctx, conn, customerId)
}

func GetHistoryByEmployeeId(ctx echo.Context, conn *gorm.DB, employeeId int) (*[]entity.HistoryReport, error) {
	return repositories.GetHistoryByEmployeeId(ctx, conn, employeeId)
}

func GetHistoryByHistoryId(ctx echo.Context, conn *gorm.DB, historyId int) (*entity.HistoryReport, error) {
	return repositories.GetHistoryByHistoryId(ctx, conn, historyId)
}
