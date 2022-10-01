package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rooms(ctx echo.Context, conn *gorm.DB) (*[]entity.Room, error) {
	rooms, err := repositories.Rooms(ctx, conn)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func RoomsStatus(ctx echo.Context, conn *gorm.DB, req request.RoomsStatus) (string, error) {
	timenow := pkg.GetDatetime()
	model := entity.RoomsStatus{
		RoomId:   req.RoomId,
		Status:   req.Status,
		UpdateAt: timenow,
		UpdateBy: req.UpdateBy,
	}
	err := repositories.RoomsStatus(ctx, conn, model)
	if err != nil {
		return "Can not update", err
	}
	return "Update success", nil
}

func RoomInsert(ctx echo.Context, conn *gorm.DB, req request.RoomInsert) (string, error) {
	timenow := pkg.GetDatetime()
	req.Status = "I"
	model := entity.RoomInsert{
		RoomNum:     req.RoomNum,
		Floors:      req.Floors,
		Description: req.Description,
		BuildingId:  req.BuildingId,
		Status:      req.Status,
		UpdateAt:    timenow,
		UpdateBy:    req.UpdateBy,
		CreateAt:    timenow,
	}
	err := repositories.RoomInsert(ctx, conn, model)
	if err != nil {
		return "Can not insert", err
	}
	return "Insert success", nil
}

func RoomByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Room, error) {
	res, err := repositories.RoomByDormId(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
