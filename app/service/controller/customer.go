package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
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

func CustomerViewProfile(ctx echo.Context, conn *gorm.DB, email string) (*response.CustomerProfile, error) {
	req := request.CustomerProfile{
		Email: email,
	}
	res, err := service.CustomerViewProfile(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CustomerEditProfile(ctx echo.Context, conn *gorm.DB, req request.CustomerEditProfile, email string) error {
	err := service.CustomerEditProfile(ctx, conn, req, email)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerById(ctx echo.Context, conn *gorm.DB, customerId int64) (*entity.Customer, error) {
	id := int(customerId)
	cus, err := service.GetCustomerById(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return cus, nil
}
