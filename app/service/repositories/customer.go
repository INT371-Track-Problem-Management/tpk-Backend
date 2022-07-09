package repositories

import (
	entity "tpk-backend/app/model/entity"

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

func GetUserByCustomerId(ctx echo.Context, conn *gorm.DB, customerId int) (*entity.Customer, error) {
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
