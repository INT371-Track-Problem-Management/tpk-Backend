package service

import (
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestService(ctx echo.Context, req request.Test, conn *gorm.DB) (*response.Test, error) {
	data, err := repositories.TestRepository(ctx, req, conn)
	if err != nil {
		return nil, err
	}
	res := &response.Test{
		UserId: data.UserId,
		Name:   data.Name,
	}
	return res, nil
}
