package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FetcStatDashBoard(ctx echo.Context, conn *gorm.DB, req request.Stat) (*entity.Stat, error) {
	stat, err := service.FetcStatDashBoard(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return stat, nil
}
