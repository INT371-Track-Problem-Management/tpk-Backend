package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Customer(ctx echo.Context, conn *gorm.DB) (*[]entity.Customer, error) {
	Customer, err := repositories.Customer(ctx, conn)
	if err != nil {
		return nil, err
	}
	return Customer, nil
}

func CustomerViewProfile(ctx echo.Context, conn *gorm.DB, req request.CustomerProfile) (*response.CustomerProfile, error) {
	cus, err := repositories.CustomerViewProfile(ctx, conn, req)
	if err != nil {
		return nil, err
	}
	profile := response.CustomerProfile{
		CustomerId:  cus.CustomerId,
		Email:       cus.Email,
		Fname:       cus.Fname,
		Lname:       cus.Lname,
		Sex:         cus.Sex,
		DateOfBirth: cus.DateOfBirth,
		Phone:       cus.DateOfBirth,
		Age:         cus.Age,
	}
	return &profile, nil
}

func CustomerEditProfile(ctx echo.Context, conn *gorm.DB, req request.CustomerEditProfile, email string) error {
	err := repositories.CustomerEditProfile(ctx, conn, req, email)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerById(ctx echo.Context, conn *gorm.DB, customerId int) (*entity.Customer, error) {
	cus, err := repositories.GetCustomerById(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	profile := entity.Customer{
		CustomerId:  cus.CustomerId,
		Email:       cus.Email,
		Fname:       cus.Fname,
		Lname:       cus.Lname,
		Sex:         cus.Sex,
		DateOfBirth: cus.DateOfBirth,
		Phone:       cus.Phone,
		Age:         cus.Age,
		Address:     cus.Address,
		Status:      cus.Status,
	}
	return &profile, nil
}
