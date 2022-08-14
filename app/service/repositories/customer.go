package repositories

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Customer(ctx echo.Context, conn *gorm.DB) (*[]entity.Customer, error) {
	var Customer []entity.Customer
	err := conn.Table("customer").Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}

func CustomerViewProfile(ctx echo.Context, conn *gorm.DB, req request.CustomerProfile) (*entity.Customer, error) {
	var Customer entity.Customer
	err := conn.Table("customer").Where("email = ?", req.Email).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}

func GetCustomerById(ctx echo.Context, conn *gorm.DB, customerId int) (*entity.Customer, error) {
	var Customer *entity.Customer
	err := conn.Table("customer").Where("customerId = ?", customerId).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return Customer, nil
}

func CustomerByEmail(ctx echo.Context, conn *gorm.DB, email string) (*entity.Customer, error) {
	var Customer entity.Customer
	err := conn.Table("customer").Where("email = ?", email).Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}

func CustomerEditProfile(ctx echo.Context, conn *gorm.DB, req request.CustomerEditProfile, email string) error {
	err := conn.Table("customer").Where("email = ?", email).Updates(req).Error
	if err != nil {
		return err
	}
	return nil
}
