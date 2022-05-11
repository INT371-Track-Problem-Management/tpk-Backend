package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rooms(ctx echo.Context, conn *gorm.DB) (*[]entity.Room, error) {
	res, err := service.Rooms(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
