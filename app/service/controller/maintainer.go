package controller

import (
	"tpk-backend/app/model/entity"
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

func Maintainerlist(ctx echo.Context, conn *gorm.DB) ([]*entity.Maintainer, error) {
	maintainers, err := service.Maintainerlist(ctx, conn)
	if err != nil {
		return nil, err
	}
	return maintainers, nil
}

func MaintainerById(ctx echo.Context, conn *gorm.DB, maintainerId string) (*entity.Maintainer, error) {
	maintainer, err := service.MaintainerById(ctx, conn, maintainerId)
	if err != nil {
		return nil, err
	}
	return maintainer, nil
}
