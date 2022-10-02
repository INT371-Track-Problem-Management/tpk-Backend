package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddMaintainer(ctx echo.Context, conn *gorm.DB, req request.Maintainer) (*int, error) {
	id, err := service.AddMaintainer(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return id, nil
}
