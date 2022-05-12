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
