package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, req request.RoomAddCustomer) error {
	err := repositories.RoomAddCustomer(ctx, conn, req)
	if err != nil {
		return err
	}
	return nil
}

func RoomRemoveCustomer(ctx echo.Context, conn *gorm.DB, id int) error {
	err := repositories.RoomRemoveCustomer(ctx, conn, id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllRoomWithCustomer(ctx echo.Context, conn *gorm.DB, dormId int) ([]*entity.RoomJoinDorm, error) {
	res, err := repositories.GetAllRoomWithCustomer(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
