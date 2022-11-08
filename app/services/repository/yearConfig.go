package repository

import (
	"tpk-backend/app/models/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func YearConfig(ctx echo.Context, conn *gorm.DB) (*response.Year, error) {
	var year response.Year
	err := conn.Table("yearConfig").Order("year desc").Select("year").Scan(&year.Year).Error
	if err != nil {
		return nil, err
	}
	return &year, nil
}
