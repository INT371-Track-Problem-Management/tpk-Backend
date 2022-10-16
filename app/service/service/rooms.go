package service

import (
	"fmt"
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"
	"tpk-backend/app/model/response"
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

func RoomInsert(ctx echo.Context, conn *gorm.DB, req request.RoomInsert) error {
	session := conn.Begin()
	timenow := pkg.GetDatetime()
	for _, room := range req.Rooms {
		model := entity.RoomInsert{
			RoomNum:     room.RoomNum,
			Floors:      room.Floors,
			Description: room.Description,
			BuildingId:  req.BuildingId,
			Status:      "I",
			UpdateAt:    timenow,
			UpdateBy:    req.UpdateBy,
			CreateAt:    timenow,
		}
		err := repositories.RoomInsert(ctx, session, model)
		if err != nil {
			return err
		}
	}
	session.Commit()
	return nil
}

func RoomByBuildingId(ctx echo.Context, conn *gorm.DB, buildingId string) (*response.RoomByBuildingId, error) {
	totalFloor, err := repositories.TotalFlooorsByBuildingId(ctx, conn, buildingId)
	if err != nil {
		return nil, err
	}
	rooms := []interface{}{}
	for i := 1; i < *totalFloor+1; i++ {
		floorNum := fmt.Sprintf(`floor%v`, i)
		allroomfloor, err := repositories.RoomInFloorByBuildingId(ctx, conn, buildingId, i)
		if err != nil {
			return nil, err
		}
		room := map[string]interface{}{
			floorNum: allroomfloor,
		}
		rooms = append(rooms, room)
	}

	allroom := response.RoomByBuildingId{
		BuildingId: buildingId,
		Floors:     rooms,
	}

	return &allroom, nil
}

func RoomByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*entity.Room, error) {
	room, err := repositories.RoomByRoomId(ctx, conn, roomId)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func RoomByRoomNum(ctx echo.Context, conn *gorm.DB, roomNum string) (*entity.Room, error) {
	room, err := repositories.RoomByRoomNum(ctx, conn, roomNum)
	if err != nil {
		return nil, err
	}
	return room, nil
}
