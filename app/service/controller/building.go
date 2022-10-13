package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Building(ctx echo.Context, conn *gorm.DB, req request.Building) (*response.Building, error) {
	res, err := service.Building(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func BuildingInsert(ctx echo.Context, conn *gorm.DB, req request.BuildingInsert) (*string, error) {
	res, err := service.BuildingInsert(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func BuildingDelete(ctx echo.Context, conn *gorm.DB, req request.BuildingDelete) (*string, error) {
	res, err := service.BuildingDelete(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
