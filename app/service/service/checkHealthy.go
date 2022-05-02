package service

import (
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CheckHealthy(ctx echo.Context, conn *gorm.DB) (*string, error) {
	data, err := repositories.CheckHealthy(ctx, conn)
	if err != nil {
		return nil, err
	}
	return data, nil
}
