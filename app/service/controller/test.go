package controller

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestController(ctx echo.Context, req request.Test, conn *gorm.DB) (*response.Test, error) {
	res, err := service.TestService(ctx, req, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}
