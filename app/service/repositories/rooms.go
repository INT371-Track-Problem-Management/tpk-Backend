package repositories

import (
	"fmt"
	entity "tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rooms(ctx echo.Context, conn *gorm.DB) (*[]entity.Room, error) {
	var room []entity.Room
	err := conn.Table("room").Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func RoomsStatus(ctx echo.Context, conn *gorm.DB, req entity.RoomsStatus) error {
	err := conn.Exec("UPDATE room SET status = ?, updateAt = ?, updateBy = ? WHERE roomId = ?", req.Status, req.UpdateAt, req.UpdateBy, req.RoomId).Error
	if err != nil {
		return err
	}
	return nil
}

func RoomInsert(ctx echo.Context, conn *gorm.DB, model entity.RoomInsert) error {
	err := conn.Table("room").Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func RoomByBuildingId(ctx echo.Context, conn *gorm.DB, buildingId string) (*[]entity.Room, error) {
	var room []entity.Room
	err := conn.Table("room").Where("buildingId = ?", buildingId).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func TotalFlooorsByBuildingId(ctx echo.Context, conn *gorm.DB, buildingId string) (*int, error) {
	var floors int
	sql := fmt.Sprintf(`
		SELECT MAX(r.floors)
		FROM room r 
		WHERE r.buildingId = %v;
	`, buildingId)
	err := conn.Raw(sql).Scan(&floors).Error
	if err != nil {
		return nil, err
	}
	return &floors, nil
}

func RoomInFloorByBuildingId(ctx echo.Context, conn *gorm.DB, buildingId string, floor int) (*[]entity.RoomByFloors, error) {
	var rooms []entity.RoomByFloors
	sql := fmt.Sprintf(`
		SELECT r.*
		FROM room r 
		WHERE r.buildingId = %v
		AND r.floors = %v;
	`, buildingId, floor)
	err := conn.Raw(sql).Scan(&rooms).Error
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}

func RoomByRoomId(ctx echo.Context, conn *gorm.DB, roomId string) (*entity.Room, error) {
	var room entity.Room
	err := conn.Table("room").Where("roomId = ?", roomId).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func RoomByRoomNum(ctx echo.Context, conn *gorm.DB, roomNum string) (*entity.Room, error) {
	var room entity.Room
	err := conn.Table("room").Where("roomNum = ?", roomNum).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}
