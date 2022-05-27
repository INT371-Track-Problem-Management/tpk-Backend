package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Customer(ctx echo.Context, conn *gorm.DB) (*[]entity.Customer, error) {
	res, err := service.Customer(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
