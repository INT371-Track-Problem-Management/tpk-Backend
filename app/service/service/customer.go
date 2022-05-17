package service

import (
	entity "tpk-backend/app/model/entity"
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
