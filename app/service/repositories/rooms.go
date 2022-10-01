package repositories

import (
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

func RoomByDormId(ctx echo.Context, conn *gorm.DB, dormId string) (*[]entity.Room, error) {
	var room []entity.Room
	err := conn.Table("room").Where("dormId = ?", dormId).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}
