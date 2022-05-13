package service

import (
	"fmt"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dorm(ctx echo.Context, conn *gorm.DB, req request.Dorm) (*response.Dorm, error) {
	data, err := repositories.Dorm(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	res := &response.Dorm{
		DormId:  data.DormId,
		Address: data.Address,
		Phone:   data.Phone,
		Email:   data.Email,
	}
	return res, nil
}
