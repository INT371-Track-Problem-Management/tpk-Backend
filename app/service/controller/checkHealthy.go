package controller

import (
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CheckHealthy(ctx echo.Context, conn *gorm.DB) (*string, error) {
	res, err := service.CheckHealthy(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
