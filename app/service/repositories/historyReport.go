package repositories

import (
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetHistoryByCustomerId(ctx echo.Context, conn gorm.DB, customerId string) (*entity.HistoryReport, error) {
	history := new(entity.HistoryReport)
	err := conn.Table("historyReport").Where("customerId = ?", customerId).Scan(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}
