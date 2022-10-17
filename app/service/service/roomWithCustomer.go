package service

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/pkg"
	"tpk-backend/app/service/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoomAddCustomer(ctx echo.Context, conn *gorm.DB, req request.RoomAddCustomer) error {
	timenow := pkg.GetDatetime()
	model := entity.RoomAddCustomer{
		RoomId:     req.RoomId,
		CustomerId: req.CustomerId,
		BuildingId: req.BuildingId,
		Status:     "A",
		CreateAt:   timenow,
		UpdateAt:   timenow,
		UpdateBy:   req.UpdateBy,
	}
	err := repositories.RoomAddCustomer(ctx, conn, model)
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
func GetAllRoomWithCustomer(ctx echo.Context, conn *gorm.DB, dormId int) ([]*entity.RoomJoinBulding, error) {
	res, err := repositories.GetAllRoomWithCustomer(ctx, conn, dormId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetRoomWithCustomerId(ctx echo.Context, conn *gorm.DB, customerId string) (*[]entity.RoomWithCustomerId, error) {
	rooms, err := repositories.GetRoomWithCustomerId(ctx, conn, customerId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
