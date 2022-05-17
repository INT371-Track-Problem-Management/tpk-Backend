package repositories

import (
	entity "tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Customer(ctx echo.Context, conn *gorm.DB) (*[]entity.Customer, error) {
	var Customer []entity.Customer
	err := conn.Table("room").Find(&Customer).Error
	if err != nil {
		return nil, err
	}
	return &Customer, nil
}
