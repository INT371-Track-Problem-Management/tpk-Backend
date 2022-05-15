package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dorm(ctx echo.Context, conn *gorm.DB, req request.Dorm) (*response.Dorm, error) {
	res, err := service.Dorm(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DormInsert(ctx echo.Context, conn *gorm.DB, req request.DormInsert) (*string, error) {
	res, err := service.DormInsert(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DormDelete(ctx echo.Context, conn *gorm.DB, req request.DormDelete) (*string, error) {
	res, err := service.DormDelete(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
