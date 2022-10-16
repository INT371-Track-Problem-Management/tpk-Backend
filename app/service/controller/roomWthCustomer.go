package controller

import (
	"net/http"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, req request.RoomAddCustomer) error {
	err := service.RoomAddCustomer(ctx, conn, req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func RoomRemoveCustomer(ctx echo.Context, conn *gorm.DB, id int64) error {
	rwcId := int(id)
	err := service.RoomRemoveCustomer(ctx, conn, rwcId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, "success")
}

func GetAllRoomWithCustomer(ctx echo.Context, conn *gorm.DB, dormId int64) ([]*entity.RoomJoinBulding, error) {
	id := int(dormId)
	res, err := service.GetAllRoomWithCustomer(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetRoomWithCustomerId(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.RoomWithCustomerId, error) {
	rooms, err := service.GetRoomWithCustomerId(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
