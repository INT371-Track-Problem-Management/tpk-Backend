package repositories

import (
	entity "tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

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

func RoomsStatus(ctx echo.Context, conn *gorm.DB, req request.RoomsStatus) error {
	err := conn.Exec("UPDATE room SET status = ? WHERE roomId = ?", req.Status, req.RoomId).Error
	if err != nil {
		return err
	}
	return nil
}
