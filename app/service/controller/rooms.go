package controller

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
	"tpk-backend/app/service/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rooms(ctx echo.Context, conn *gorm.DB) (*[]entity.Room, error) {
	res, err := service.Rooms(ctx, conn)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func RoomsStatus(ctx echo.Context, conn *gorm.DB, req request.RoomsStatus) (string, error) {
	res, err := service.RoomsStatus(ctx, conn, req)
	if err != nil {
		return "", err
	}
	return res, nil
}

func RoomInsert(ctx echo.Context, conn *gorm.DB, req request.RoomInsert) error {
	err := service.RoomInsert(ctx, conn, req)
	if err != nil {
		return err
	}
	return nil
}

func RoomByBuildingId(ctx echo.Context, conn *gorm.DB, dormId string) (*response.RoomByBuildingId, error) {
	res, err := service.RoomByBuildingId(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func RoomByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*entity.Room, error) {
	room, err := service.RoomByRoomId(ctx, conn, roomId)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func RoomByRoomNum(ctx echo.Context, conn *gorm.DB, roomNum string) (*entity.Room, error) {
	room, err := service.RoomByRoomNum(ctx, conn, roomNum)
	if err != nil {
		return nil, err
	}
	return room, nil
}
