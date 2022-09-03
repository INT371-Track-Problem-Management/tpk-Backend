package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetHistoryByCustomerId(ctx echo.Context, conn *gorm.DB, customerId int) (*[]entity.HistoryReport, error) {
	history := new([]entity.HistoryReport)
	err := conn.Table("historyReport").Where("customerId = ?", customerId).Scan(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

func GetHistoryByEmployeeId(ctx echo.Context, conn *gorm.DB, employeeId int) (*[]entity.HistoryReport, error) {
	history := new([]entity.HistoryReport)
	err := conn.Table("historyReport").Where("employeeId = ?", employeeId).Scan(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

func GetHistoryByHistoryId(ctx echo.Context, conn *gorm.DB, historyId int) (*entity.HistoryReport, error) {
	history := new(entity.HistoryReport)
	err := conn.Table("historyReport").Where("historyId = ?", historyId).Scan(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}
